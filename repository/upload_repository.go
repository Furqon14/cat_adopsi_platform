package repository

import (
	"cat_adoption_platform/model"
	"database/sql"
)

type UploadImageRepository interface {
	Save(image model.Image) error
}

type uploadImageRepository struct {
	db     *sql.DB
	images []model.Image
}

func (r *uploadImageRepository) Save(image model.Image) error {
	r.images = append(r.images, image)
	return nil
}

func NewUploadImageRepository(db *sql.DB) UploadImageRepository {
	return &uploadImageRepository{db: db}
}
