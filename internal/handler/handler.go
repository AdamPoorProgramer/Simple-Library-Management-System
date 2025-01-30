package handler

import (
	"LIBRARY-API-SERVER/api/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

type Model interface {
	TableName() string
}
type Handler[T Model] struct {
	*zap.Logger
	db *gorm.DB
}

func NewHandler[T Model](db *gorm.DB, log *zap.Logger) *Handler[T] {
	return &Handler[T]{
		db:     db,
		Logger: log,
	}
}
func PreLoad(db *gorm.DB, modelInstance interface{}) *gorm.DB {
	switch modelInstance.(type) {
	case *model.Book:
		return db.Preload("Category") // برای مدل Book، Category را پری‌لود می‌کنیم
	case *model.Borrowing:
		return db.Preload("Member").Preload("Book").Preload("Book.Category") // برای Borrowing، Member و Book.Category پری‌لود می‌شود
	case *model.Category:
		return db.Preload("Book") // برای Category، Book پری‌لود می‌شود
	default:
		return db
	}
}

func (h Handler[T]) Post(c *gin.Context) {
	var modelInstance T
	if err := c.ShouldBindJSON(&modelInstance); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		h.Logger.Error("Error occurred while binding JSON.", zap.Error(err))
		return
	}
	if err := h.db.Create(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to create " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while creating record.", zap.Error(err))
		return
	}
	// Reload the record with preloaded relationships
	if err := PreLoad(h.db, &modelInstance).First(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to preload " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while preloading record.", zap.Error(err))
		return
	}
	c.JSON(200, modelInstance)
	h.Logger.Info("Record created successfully.", zap.String(modelInstance.TableName(), fmt.Sprintf("%+v", modelInstance)))
}
func (h Handler[T]) GetAll(c *gin.Context) {
	var models []T
	var modelInstance T
	if err := h.db.Find(&models).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to get records"})
		h.Logger.Error("Error occurred while getting records.", zap.Error(err))
		return
	}
	if err := PreLoad(h.db, &modelInstance).Find(&models).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to preload " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while preloading record.", zap.Error(err))
		return
	}
	c.JSON(200, models)
	h.Logger.Info("Records retrieved successfully.")
}
func (h Handler[T]) GetById(c *gin.Context) {
	if c.Query("id") == "" {
		h.GetAll(c)
		return
	}

	var modelInstance T
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		h.Logger.Error("Error occurred while parsing ID.", zap.Error(err))
		return
	}
	if err := h.db.Where("ID = ?", id).First(&modelInstance).Error; err != nil {
		c.JSON(404, gin.H{"error": modelInstance.TableName() + " not found"})
		h.Logger.Error("Error occurred while getting record by ID.", zap.Error(err))
		return
	}
	if err := PreLoad(h.db, &modelInstance).First(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to preload " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while preloading record.", zap.Error(err))
		return
	}
	c.JSON(200, modelInstance)
	h.Logger.Info("Record retrieved successfully by ID.", zap.String(modelInstance.TableName(), fmt.Sprintf("%+v", modelInstance)))
}
func (h Handler[T]) Put(c *gin.Context) {
	var modelInstance T
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		h.Logger.Error("Error occurred while parsing ID.", zap.Error(err))
		return
	}

	// دریافت داده‌های جدید از کاربر
	if err := c.ShouldBindJSON(&modelInstance); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		h.Logger.Error("Error occurred while binding JSON.", zap.Error(err))
		return
	}

	// **1. دریافت داده‌ی قبلی از دیتابیس**
	var existingModelInstance T
	if err := h.db.Where("ID = ?", id).First(&existingModelInstance).Error; err != nil {
		c.JSON(404, gin.H{"error": modelInstance.TableName() + " not found"})
		h.Logger.Error("Error occurred while getting record by ID.", zap.Error(err))
		return
	}

	// **2. بررسی و مدیریت روابط many-to-many**
	switch v := any(&modelInstance).(type) {
	case *model.Borrowing:
		// حذف روابط قدیمی و جایگزینی با جدید
		h.db.Model(&v).Association("Member").Replace(v.Member)
		h.db.Model(&v).Association("Book").Replace(v.Book)
	case *model.Book:
		h.db.Model(&v).Association("Category").Replace(v.Category)
	case *model.Category:
		h.db.Model(&v).Association("Book").Replace(v.Book)
	}

	// **3. اعمال آپدیت روی سایر فیلدها (به جز many-to-many)**
	if err := h.db.Model(&existingModelInstance).Where("ID = ?", id).Updates(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to update " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while updating record by ID.", zap.Error(err))
		return
	}

	// **4. پریلود داده‌ها پس از آپدیت**
	if err := PreLoad(h.db, &modelInstance).First(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to preload " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while preloading record.", zap.Error(err))
		return
	}

	// **5. ارسال پاسخ به کاربر**
	c.JSON(200, modelInstance)
	h.Logger.Info("Record updated successfully by ID.", zap.String(modelInstance.TableName(), fmt.Sprintf("%+v", modelInstance)))
}

func (h Handler[T]) Delete(c *gin.Context) {
	var modelInstance T
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		h.Logger.Error("Error occurred while parsing ID.", zap.Error(err))
		return
	}
	if err := h.db.Where("ID =?", id).Delete(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to delete " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while deleting record by ID.", zap.Error(err))
		return
	}
	c.JSON(200, gin.H{"message": modelInstance.TableName() + " has been deleted"})
	h.Logger.Info("Record deleted successfully by ID.", zap.String(modelInstance.TableName(), fmt.Sprintf("%+v", modelInstance)))
}

func (h Handler[T]) Register(router *gin.RouterGroup) {
	var modelInstance T
	tableName := modelInstance.TableName()
	router.GET("", h.GetById)
	router.POST("", h.Post)
	router.PUT("", h.Put)
	router.DELETE("", h.Delete)
	h.Logger.Info("Routes registered for " + tableName)
}
