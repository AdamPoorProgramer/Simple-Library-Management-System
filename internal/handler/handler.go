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
	s := fmt.Sprintf("%s created successfully", modelInstance.TableName())
	c.JSON(200, gin.H{"message": s, modelInstance.TableName(): modelInstance})
	h.Logger.Info("Record created successfully.", zap.String(modelInstance.TableName(), fmt.Sprintf("%+v", modelInstance)))
}
func (h Handler[T]) GetAllMembers(c *gin.Context) {
	var models []T
	var modelInstance T
	if err := h.db.Find(&models).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to get records"})
		h.Logger.Error("Error occurred while getting records.", zap.Error(err))
		return
	}
	c.JSON(200, gin.H{modelInstance.TableName(): models})
	h.Logger.Info("Records retrieved successfully.")
}
func (h Handler[T]) GetById(c *gin.Context) {
	var modelInstance T
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
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
	c.JSON(200, gin.H{modelInstance.TableName(): modelInstance})
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
	if err := c.ShouldBindJSON(&modelInstance); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		h.Logger.Error("Error occurred while binding JSON.", zap.Error(err))
		return
	}
	if err := h.db.Model(&model.Member{}).Where("ID = ?", id).Updates(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to update " + modelInstance.TableName()})
		h.Logger.Error("Error occurred while updating record by ID.", zap.Error(err))
		return
	}
	c.JSON(200, gin.H{"message": modelInstance.TableName() + " hs been updated", modelInstance.TableName(): modelInstance})
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
	router.GET("/:id", h.GetById)
	router.GET("", h.GetAllMembers)
	router.POST("", h.Post)
	router.PUT("/", h.Put)
	router.DELETE("/", h.Delete)
	h.Logger.Info("Routes registered for " + tableName)
}
