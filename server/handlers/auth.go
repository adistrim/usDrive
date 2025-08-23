package handlers

import (
	"log"
	"net/http"
	"usdrive/services"

	"github.com/gin-gonic/gin"
)

type SignInRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}

func GoogleSignInHandler(c *gin.Context) {
	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	response, err := services.SignInWithGoogle(c.Request.Context(), req.IDToken)
	if err != nil {
		log.Printf("Error during sign-in: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	
	c.JSON(http.StatusOK, response)
}
