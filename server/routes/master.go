package routes

import (
	"github.com/gin-gonic/gin"
	
	"usdrive/handlers"
)

func Master() *gin.Engine {
	router := gin.Default()
	
	router.GET("/health", HealthCheck)
	
	router.POST("/api/request/upload", handlers.RequestUpload)
	router.POST("/api/files/:fileId/complete", handlers.CompleteUpload)
	
	return router
}
