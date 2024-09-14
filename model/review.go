package model

import (
	"time"
)

type Review struct {
	ReviewID  string    `json:"review_id"`
	UserID    string    `json:"user_id"`
	CatID     string    `json:"cat_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
