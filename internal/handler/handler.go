package handler

import (
	"LIBRARY-API-SERVER/api/model"
	"LIBRARY-API-SERVER/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Model interface {
	TableName() string
}
type Handler[T Model] struct {
	*logger.Logger
	db *gorm.DB
}

func NewHandler[T Model](db *gorm.DB, log *logger.Logger) Handler[T] {
	return Handler[T]{
		db:     db,
		Logger: log,
	}
}

func (h Handler[T]) Post(c *gin.Context) {
	var modelInstance T
	if err := c.ShouldBindJSON(&modelInstance); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.db.Create(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to create " + modelInstance.TableName()})
		return
	}
	s := fmt.Sprintf("%s created successfully", modelInstance.TableName())
	c.JSON(200, gin.H{"message": s, modelInstance.TableName(): modelInstance})
}
func (h Handler[T]) GetAllMembers(c *gin.Context) {
	var models []T
	if err := h.db.Find(&models).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to get " + T.TableName(nil)})
		return
	}
	c.JSON(200, gin.H{T.TableName(nil): models})
}
func (h Handler[T]) GetById(c *gin.Context) {
	var modelInstance T
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.db.Where("ID = ?", id).First(&modelInstance).Error; err != nil {
		c.JSON(404, gin.H{"error": T.TableName(nil) + " not found"})
		return
	}
	c.JSON(200, gin.H{T.TableName(nil): modelInstance})
}
func (h Handler[T]) Put(c *gin.Context) {
	var modelInstance T
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.ShouldBindJSON(&modelInstance); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.db.Model(&model.Member{}).Where("ID = ?", id).Updates(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to update " + T.TableName(nil)})
		return
	}
	c.JSON(200, gin.H{"message": modelInstance.TableName() + " hs been updated", modelInstance.TableName(): modelInstance})
}
func (h Handler[T]) Delete(c *gin.Context) {
	var modelInstance T
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.db.Where("ID =?", id).Delete(&modelInstance).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to delete " + modelInstance.TableName()})
		return
	}
	c.JSON(200, gin.H{"message": modelInstance.TableName() + "has been deleted"})
}

func (h Handler[T]) Register(router *gin.RouterGroup, log *logger.Logger) {
	router.GET("/"+T.TableName(nil)+"/:id", h.GetById)
	router.GET("/"+T.TableName(nil), h.GetAllMembers)
	router.POST("/"+T.TableName(nil), h.Post)
	router.PUT("/"+T.TableName(nil)+"/:id", h.Put)
	router.DELETE("/"+T.TableName(nil)+"/:id", h.Delete)
}
