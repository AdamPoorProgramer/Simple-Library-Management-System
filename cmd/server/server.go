package main

import (
	"LIBRARY-API-SERVER/configs"
	"LIBRARY-API-SERVER/internal/db"
	"LIBRARY-API-SERVER/internal/db/sqlite"
	"LIBRARY-API-SERVER/internal/router"
	"LIBRARY-API-SERVER/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	log, err := logger.NewLogger(zap.DebugLevel)
	if err != nil {
		fmt.Print(err.Error())
	}
	server := gin.Default()
	dataBase := sqlite.NewSQLiteOrPanic()
	err = db.Migrate(dataBase)
	if err != nil {
		panic(err)
	}
	router.SetupRoutes(dataBase, server, log)
	server.Run(config.Server.Host + ":" + config.Server.Port)
}
