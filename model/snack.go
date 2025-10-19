package model

type Snack struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Unit     int     `json:"unit"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}
