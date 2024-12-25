package main

import (
	"LIBRARY-API-SERVER/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Run(configs.LoadConfig().Server.URL)
}
