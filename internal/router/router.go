package router

import (
	"LIBRARY-API-SERVER/api/model"
	"LIBRARY-API-SERVER/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, router *gin.Engine) {
	apiRouter := router.Group("/api")
	handler.NewHandler[model.Book](db).Register(apiRouter)
	handler.NewHandler[model.Member](db).Register(apiRouter)
	handler.NewHandler[model.Borrowing](db).Register(apiRouter)
	handler.NewHandler[model.Category](db).Register(apiRouter)
	handler.NewHandler[model.BookCategory](db).Register(apiRouter)
}
