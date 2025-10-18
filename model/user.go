package model

import "database/sql"

type User struct {
	ID             int            `json:"id"`
	Role           string         `json:"role"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	PhoneNumber    string         `json:"phone_number"`
	Language       string         `json:"language"`
	Status         string         `json:"status_user"`
	ProfilePicture sql.NullString `json:"imageURL"`
	CreatedAt      string         `json:"createdAt"`
	UpdatedAt      string         `json:"updatedAt"`
}

type ProfileUser struct {
	ID             int     `json:"id"`
	Role           string  `json:"role"`
	Username       string  `json:"username"`
	Email          string  `json:"email"`
	PhoneNumber    string  `json:"phone_number"`
	Language       string  `json:"language"`
	Status         string  `json:"status_user"`
	ProfilePicture *string `json:"imageURL"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
