package main

import (
	"laakso-go-backend/internal/config"
	"laakso-go-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/visitorCounter", handlers.GetVisitorCounter)
	r.POST("/visitorCounter", handlers.IncrementVisitorCounter)

	r.Run(":8080")
}
