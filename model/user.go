package model

import (
	"database/sql"
)

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

type UserDTO struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
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

type UserUpdateRequest struct {
	// 🔥 Semua field yang mungkin diubah harus menggunakan pointer
	ID             *int    `json:"id"`
	Username       *string `json:"name"`
	Email          *string `json:"email"`
	ProfilePicture *string `json:"imageURL"`
	Password       *string `json:"password"`
	PhoneNumber    *string `json:"phone_number"`
	Language       *string `json:"language"`
	Role           *string `json:"role"`
	Status         *string `json:"status"`
	CreatedAt      *string `json:"createdAt"`
	UpdatedAt      *string `json:"updatedAt"`
}

// UserResponse (untuk respons setelah update)
type UserResponse struct {
	ID             int     `json:"id"`
	Username       string  `json:"name"`
	Email          string  `json:"email"`
	ProfilePicture *string `json:"imageURL"`
	Password       *string `json:"password"`
	PhoneNumber    string  `json:"phone_number"`
	Language       string  `json:"language"`
	Role           string  `json:"role"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
}

// UploadResponse untuk POST /uploads
type UploadResponse struct {
	ImageURL string `json:"imageURL"`
}
