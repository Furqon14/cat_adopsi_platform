package repository

import (
	"database/sql"
)

type MidtransRepository interface {
	Payment() error
}

type midtransRepository struct {
	db *sql.DB
}

func (r *midtransRepository) Payment() error {
	return nil
}

func NewMidtransRepository(db *sql.DB) MidtransRepository {
	return &midtransRepository{db: db}
}
