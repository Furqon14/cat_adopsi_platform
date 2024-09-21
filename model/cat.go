package model

import (
	"mime/multipart"
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

type FormUploadImages struct {
	Image     multipart.FileHeader `form:"file"`
	ImageData CatImage             `form:"image_data"`
}

type CatImage struct {
	ID        string `json:"id"`
	CatID     string `json:"cat_id"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
