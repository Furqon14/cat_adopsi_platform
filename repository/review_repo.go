package repository

import (
	"cat_adoption_platform/model"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ReviewRepository interface {
	Create(payload model.Review) (model.Review, error)
	GetByID(reviewId uuid.UUID) (model.Review, error)
	Update(payload model.Review) (model.Review, error)
	Delete(id uuid.UUID) error
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
	fmt.Println(payload)

	query := `
        INSERT INTO t_review (user_id, cat_id, rating, comment)
        VALUES ($1, $2, $3, $4)
        RETURNING review_id, user_id, cat_id, rating, comment, created_at, updated_at
    `

	err := r.db.QueryRow(query, payload.UserID, payload.CatID, payload.Rating, payload.Comment).Scan(
		&review.ReviewID, &review.UserID, &review.CatID, &review.Rating, &review.Comment, &review.CreatedAt, &review.UpdatedAt,
	)

	if err != nil {
		return model.Review{}, err
	}
	return review, nil
}

func (r *reviewRepository) GetByID(reviewId uuid.UUID) (model.Review, error) {
	var review model.Review

	query := `
        SELECT review_id, user_id, cat_id, rating, comment, created_at, updated_at
        FROM t_review
        WHERE review_id = $1
    `

	err := r.db.QueryRow(query, reviewId).Scan(
		&review.ReviewID, &review.UserID, &review.CatID, &review.Rating, &review.Comment, &review.CreatedAt, &review.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return model.Review{}, fmt.Errorf("review not found")
	} else if err != nil {
		return model.Review{}, err
	}

	return review, nil
}

func (r *reviewRepository) Update(payload model.Review) (model.Review, error) {
	var review model.Review

	query := `
        UPDATE t_review
        SET rating = $1, comment = $2, updated_at = $3
        WHERE review_id = $4
        RETURNING review_id, user_id, cat_id, rating, comment, created_at, updated_at
    `

	err := r.db.QueryRow(query, payload.Rating, payload.Comment, time.Now(), payload.ReviewID).Scan(&review.ReviewID, &review.UserID, &review.CatID, &review.Rating, &review.Comment, &review.CreatedAt, &review.UpdatedAt)
	if err == sql.ErrNoRows {
		return model.Review{}, fmt.Errorf("review not found")
	} else if err != nil {
		return model.Review{}, err
	}
	return review, nil
}

func (r *reviewRepository) Delete(reviewId uuid.UUID) error {
	query := `
        DELETE FROM t_review
        WHERE review_id = $1
    `

	result, err := r.db.Exec(query, reviewId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("review not found")
	}

	return nil
}

func (r *reviewRepository) GetAll() ([]model.Review, error) {
	var reviews []model.Review

	rows, err := r.db.Query(
		"SELECT review_id, user_id, cat_id, rating, comment, created_at, updated_at FROM t_review",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	totalRows := 0

	err = r.db.QueryRow("SELECT COUNT(*) FROM t_review").Scan(&totalRows)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var review model.Review
		err := rows.Scan(&review.ReviewID, &review.UserID, &review.CatID, &review.Rating, &review.Comment, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return []model.Review{}, err
		}
		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		return []model.Review{}, err
	}

	return reviews, nil
}
