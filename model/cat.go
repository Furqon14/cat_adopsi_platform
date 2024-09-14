package model

import (
	"time"
)

type Cat struct {
	CatID             string    `json:"cat_id"`
	Name              string    `json:"name"`
	Breed             string    `json:"breed"`
	Age               int       `json:"age"`
	Color             string    `json:"color"`
	Description       string    `json:"description"`
	Adopted           bool      `json:"adopted"`
	Latitude          float64   `json:"latitude"`
	Longitude         float64   `json:"longitude"`
	LocationName      string    `json:"location_name"`
	PhotoURL          string    `json:"photo_url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Gender            string    `json:"gender"`
	VaccinationStatus string    `json:"vaccination_status"`
}
