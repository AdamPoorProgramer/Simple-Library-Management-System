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
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
