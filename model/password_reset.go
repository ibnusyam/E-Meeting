package model

import "time"

type PasswordReset struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
