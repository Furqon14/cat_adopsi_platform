package model

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ReviewID  uuid.UUID `json:"review_id"`
	UserID    uuid.UUID `json:"user_id"`
	CatID     uuid.UUID `json:"cat_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
