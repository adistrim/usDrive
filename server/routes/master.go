package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"usdrive/config"
	"usdrive/handlers"
)

func Master() *gin.Engine {
	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.ENV.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	router.GET("/health", HealthCheck)
	
	api := router.Group("/api")
	{
		files := api.Group("/files")
		{
			files.POST("/request/upload", handlers.RequestUpload)
			files.POST("/:fileId/complete", handlers.CompleteUpload)
			files.GET("", handlers.ListActiveFiles)
		}
		auth := api.Group("/auth")
		{
			auth.POST("/signin", handlers.GoogleSignInHandler)
		}
	}
	
	return router
}
