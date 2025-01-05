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
	bookApiRouter := apiRouter.Group("/book")
	memberApiRouter := apiRouter.Group("/book")
	categoryApiRouter := apiRouter.Group("/book")

	handler.NewBorrowingHandler[model.Book](db, log).Register(bookApiRouter, log)
	handler.NewBorrowingHandler[model.Member](db, log).Register(memberApiRouter, log)
	handler.NewBorrowingHandler[model.Category](db, log).Register(categoryApiRouter, log)
}
