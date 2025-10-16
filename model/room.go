package model

import "time"

type Room struct {
	ID        int
	Name      string
	Capacity  int
	Price     float64
	Type      string
	ImagesUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}
