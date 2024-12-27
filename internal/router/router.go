package router

import (
	"LIBRARY-API-SERVER/api/model"
	"LIBRARY-API-SERVER/internal/handler"
	"LIBRARY-API-SERVER/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, router *gin.Engine, log *logger.Logger) {
	apiRouter := router.Group("/api")
	handler.NewHandler[model.Book](db, log).Register(apiRouter, log)
	handler.NewHandler[model.Member](db, log).Register(apiRouter, log)
	handler.NewHandler[model.Borrowing](db, log).Register(apiRouter, log)
	handler.NewHandler[model.Category](db, log).Register(apiRouter, log)
	handler.NewHandler[model.BookCategory](db, log).Register(apiRouter, log)
}
