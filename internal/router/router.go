package router

import (
	"LIBRARY-API-SERVER/api/model"
	"LIBRARY-API-SERVER/internal/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, router *gin.Engine, log *zap.Logger) {
	apiRouter := router.Group("/api")
	bookApiRouter := apiRouter.Group("/book")
	memberApiRouter := apiRouter.Group("/member")
	categoryApiRouter := apiRouter.Group("/category")
	borrowingApiRouter := apiRouter.Group("/borrowing")

	handler.NewHandler[model.Book](db, log).Register(bookApiRouter)
	handler.NewHandler[model.Member](db, log).Register(memberApiRouter)
	handler.NewHandler[model.Category](db, log).Register(categoryApiRouter)
	handler.NewHandler[model.Borrowing](db, log).Register(borrowingApiRouter)
	router.Static("/static", "./static")
	handler.NewHandler[model.Book](db, log).RegFile(router)
}
