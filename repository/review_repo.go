package repository

import (
	"cat_adoption_platform/model"
	"database/sql"
	"fmt"
)

type ReviewRepository interface {
	Create(payload model.Review) (model.Review, error)
	GetByID(reviewId string) (model.Review, error)
	Update(payload model.Review) (model.Review, error)
	Delete(id string) error
	GetAll() ([]model.Review, error)
}

type reviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) Create(payload model.Review) (model.Review, error) {
	var review model.Review

	return review, nil
}

func (r *reviewRepository) GetByID(reviewId string) (model.Review, error) {
	var review model.Review

	return review, nil
}

func (r *reviewRepository) Update(payload model.Review) (model.Review, error) {
	var review model.Review

	return review, nil
}

func (r *reviewRepository) Delete(reviewId string) error {
	fmt.Println("reviewId: ", reviewId)

	return nil
}

func (r *reviewRepository) GetAll() ([]model.Review, error) {
	var reviews []model.Review

	return reviews, nil
}
