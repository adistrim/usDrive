package services

import (
	"context"
	"errors"
	"log"
	"time"
	"usdrive/config"
	"usdrive/db"
	"usdrive/models"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/idtoken"
)

func SignInWithGoogle(ctx context.Context, idToken string) (*models.SignInResponse, error) {
	gClientID := config.ENV.GClientID

	payload, err := idtoken.Validate(ctx, idToken, gClientID)
	if err != nil {
		return nil, errors.New("invalid Google ID token")
	}

	googleID := payload.Subject
	user, err := db.FindUserByGoogleID(ctx, googleID)

	if err != nil {
		if err.Error() == "not found" {
			log.Printf("User with Google ID %s not found. Creating new user.", googleID)

			newUser := &models.User{
				GoogleID:   googleID,
				Email:      payload.Claims["email"].(string),
				FullName:   payload.Claims["name"].(string),
				AvatarURL: payload.Claims["picture"].(string),
			}

			user, err = db.CreateUser(ctx, newUser)
			if err != nil {
				return nil, errors.New("failed to create new user")
			}
		} else {
			return nil, errors.New("database error while fetching user")
		}
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"eml": user.Email,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(config.ENV.JWTSecret))
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	userResponse := &models.UserResponse{
		Email:      user.Email,
		FullName:   user.FullName,
		AvatarURL: user.AvatarURL,
	}

	return &models.SignInResponse{
		AccessToken: accessToken,
		User:        userResponse,
	}, nil
}
