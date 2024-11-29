package handlers

import (
	"laakso-go-backend/internal/config"
	"laakso-go-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GET Visitor counter
func GetVisitorCounter(c *gin.Context) {
	var visitorCounter models.VisitorCounter
	if err := config.DB.First(&visitorCounter, 1).Error; err != nil {
		c.JSON(500, gin.H{"error getting visitor count": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"count": visitorCounter.Count,
	})
}

// POST Increment visitor counter
func IncrementVisitorCounter(c *gin.Context) {
	var visitorCounter models.VisitorCounter
	config.DB.First(&visitorCounter, 1)
	config.DB.Model(&visitorCounter).Update("count", visitorCounter.Count+1)

	c.JSON(200, gin.H{
		"count": visitorCounter.Count,
	})
}
