package model

import (
	"time"
)

type Adoption struct {
	AdoptionID  string    `json:"adoption_id"`
	UserID      string    `json:"user_id"`
	CatID       string    `json:"cat_id"`
	AdoptedDate time.Time `json:"adopted_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
