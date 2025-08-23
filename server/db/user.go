package db

import (
	"context"
	"errors"
	"log"
	"usdrive/models"

	"github.com/jackc/pgx/v5"
)

func FindUserByGoogleID(ctx context.Context, googleID string) (*models.User, error) {
	pool := GetDBInstance()
	query := `SELECT id, google_id, email, full_name, avatar_url, created_at FROM users WHERE google_id = $1`

	var user models.User
	err := pool.QueryRow(ctx, query, googleID).Scan(
		&user.ID,
		&user.GoogleID,
		&user.Email,
		&user.FullName,
		&user.AvatarURL,
		&user.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("not found")
		}
		log.Printf("Error querying user by Google ID: %v", err)
		return nil, err
	}

	return &user, nil
}

func CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	pool := GetDBInstance()
	query := `
        INSERT INTO users (google_id, email, full_name, avatar_url)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `

	err := pool.QueryRow(ctx, query, user.GoogleID, user.Email, user.FullName, user.AvatarURL).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	log.Printf("Successfully created user with ID: %d", user.ID)
	return user, nil
}
