package main

import (
	"LIBRARY-API-SERVER/configs"
	"LIBRARY-API-SERVER/internal/db"
	"LIBRARY-API-SERVER/internal/db/sqlite"
	"LIBRARY-API-SERVER/internal/router"
	"LIBRARY-API-SERVER/pkg/logger"
	"bufio"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	log := logger.NewLogger(zap.DebugLevel)
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config", zap.Error(err))
	}
	gin.DefaultWriter = zap.NewStdLog(log).Writer()
	server := gin.Default()
	dataBase := sqlite.NewSQLiteOrPanic(log)
	err = db.Migrate(dataBase)
	if err != nil {
		log.Fatal("Failed to migrate database", zap.Error(err))
	}
	router.SetupRoutes(dataBase, server, log)
	done := make(chan bool)
	httpServer := &http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: server.Handler(),
	}
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server", zap.Error(err))
		}
		done <- true
	}()

	select {
	case <-done:
		log.Info("Server stopped")
		return
	case <-time.After(5 * time.Second):
		log.Info("Server started", zap.String("host", config.Server.Host), zap.String("port", config.Server.Port))
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\"exit\" for shutdown server: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			log.Info("Exiting...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := httpServer.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown:", zap.Error(err))
			}
			log.Info("Server exited.")
			break
		}
		fmt.Println("wrong command")
	}
}
