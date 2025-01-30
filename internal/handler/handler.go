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

	// **1. دریافت داده‌های جدید از درخواست**
	if err := c.ShouldBindJSON(&modelInstance); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		h.Logger.Error("Error occurred while binding JSON.", zap.Error(err))
		return
	}

	// **2. دریافت داده‌ی قبلی از دیتابیس (برای حفظ ارتباطات)**
	var existingModelInstance T
	if err := h.db.Preload("Category").Preload("Book").Where("ID = ?", id).First(&existingModelInstance).Error; err != nil {
		c.JSON(404, gin.H{"error": modelInstance.TableName() + " not found"})
		h.Logger.Error("Error occurred while getting record by ID.", zap.Error(err))
		return
	}

	// **3. بررسی مدل و مدیریت روابط many-to-many**
	switch v := any(&modelInstance).(type) {
	case *model.Borrowing:
		h.db.Model(&v).Association("Member").Replace(v.Member)
		h.db.Model(&v).Association("Book").Replace(v.Book)

	case *model.Book:
		// **پاک کردن روابط قدیمی و جایگزینی دسته‌بندی‌های جدید**
		if err := h.db.Model(&v).Association("Category").Clear(); err != nil {
			c.JSON(500, gin.H{"error": "failed to clear categories"})
			h.Logger.Error("Error occurred while clearing categories.", zap.Error(err))
			return
		}
		if err := h.db.Model(&v).Association("Category").Append(v.Category); err != nil {
			c.JSON(500, gin.H{"error": "failed to update categories"})
			h.Logger.Error("Error occurred while updating categories.", zap.Error(err))
			return
		}

	case *model.Category:
		// **پاک کردن روابط قدیمی و جایگزینی کتاب‌های جدید**
		if err := h.db.Model(&v).Association("Book").Clear(); err != nil {
			c.JSON(500, gin.H{"error": "failed to clear books"})
			h.Logger.Error("Error occurred while clearing books.", zap.Error(err))
			return
		}
		if err := h.db.Model(&v).Association("Book").Append(v.Book); err != nil {
			c.JSON(500, gin.H{"error": "failed to update books"})
			h.Logger.Error("Error occurred while updating books.", zap.Error(err))
			return
		}
	}

	// **4. آپدیت سایر فیلدهای غیر مرتبط به many-to-many**
	if err := h.db.Model(&existingModelInstance).Where("ID = ?", id).Updates(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to update " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while updating record by ID.", zap.Error(err))
		return
	}

	// **5. پریلود داده‌های جدید پس از آپدیت**
	if err := PreLoad(h.db, &modelInstance).First(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to preload " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while preloading record.", zap.Error(err))
		return
	}

	// **6. ارسال پاسخ نهایی**
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
func (h Handler[T]) Home(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.File("./static/index.html")
	h.Logger.Info("Home page accessed")
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
func (h Handler[T]) RegFile(router *gin.Engine) {
	router.GET("/book", h.Home)
	router.GET("/member", h.Home)
	router.GET("/category", h.Home)
	router.GET("/borrowing", h.Home)
	router.GET("/", h.Home)
}
