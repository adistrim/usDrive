package routes

import (
	"github.com/gin-gonic/gin"
	
	"usdrive/handlers"
)

func Master() *gin.Engine {
	router := gin.Default()
	
	router.GET("/health", HealthCheck)
	
	api := router.Group("/api")
	{
		files := api.Group("/files")
		{
			files.POST("/request/upload", handlers.RequestUpload)
			files.POST("/:fileId/complete", handlers.CompleteUpload)
			files.GET("", handlers.ListActiveFiles)
		}
	}
	
	return router
}
