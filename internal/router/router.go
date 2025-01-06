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
	borrowingApiRouter := apiRouter.Group("/book")

	handler.NewHandler[model.Book](db, log).Register(bookApiRouter)
	handler.NewHandler[model.Member](db, log).Register(memberApiRouter)
	handler.NewHandler[model.Category](db, log).Register(categoryApiRouter)
	handler.NewHandler[model.Borrowing](db, log).Register(borrowingApiRouter)
}
