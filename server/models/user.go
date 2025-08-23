package models

import "time"

type User struct {
	ID         int64     `json:"id"`
	GoogleID   string    `json:"-"`
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserResponse struct {
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
}

type SignInResponse struct {
	AccessToken string        `json:"access_token"`
	User        *UserResponse `json:"user"`
}
