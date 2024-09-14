package repository

import (
	"cat_adoption_platform/model"
	"database/sql"
	"fmt"
)

type CatRepository interface {
	GetAllCats() ([]model.Cat, error)
	GetCatByID(catID string) (*model.Cat, error)
	CreateCat(cat *model.Cat) error
	UpdateCat(cat *model.Cat) error
	DeleteCat(catID string) error
}
type catRepository struct {
	db *sql.DB
}

// GetAllCats mengambil semua data kucing
func (r *catRepository) GetAllCats() ([]model.Cat, error) {
	rows, err := r.db.Query("SELECT * FROM m_cat")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []model.Cat
	for rows.Next() {
		var cat model.Cat
		if err := rows.Scan(
			&cat.CatID, &cat.Name, &cat.Breed, &cat.Age, &cat.Color,
			&cat.Description, &cat.Adopted, &cat.Latitude, &cat.Longitude,
			&cat.LocationName, &cat.PhotoURL, &cat.CreatedAt, &cat.UpdatedAt,
			&cat.Gender, &cat.VaccinationStatus,
		); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	return cats, nil
}

// GetCatByID mengambil data kucing berdasarkan ID
func (r *catRepository) GetCatByID(catID string) (*model.Cat, error) {
	row := r.db.QueryRow("SELECT * FROM m_cat WHERE cat_id = $1", catID)
	var cat model.Cat
	if err := row.Scan(
		&cat.CatID, &cat.Name, &cat.Breed, &cat.Age, &cat.Color,
		&cat.Description, &cat.Adopted, &cat.Latitude, &cat.Longitude,
		&cat.LocationName, &cat.PhotoURL, &cat.CreatedAt, &cat.UpdatedAt,
		&cat.Gender, &cat.VaccinationStatus,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cat with ID %s not found", catID)
		}
		return nil, err
	}
	return &cat, nil
}

// CreateCat menambahkan data kucing baru
func (r *catRepository) CreateCat(cat *model.Cat) error {
	_, err := r.db.Exec(
		`INSERT INTO m_cat (cat_id, name, breed, age, color, description, adopted, latitude, longitude, location_name, photo_url, created_at, updated_at, gender, vaccination_status)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
		cat.CatID, cat.Name, cat.Breed, cat.Age, cat.Color, cat.Description,
		cat.Adopted, cat.Latitude, cat.Longitude, cat.LocationName, cat.PhotoURL,
		cat.CreatedAt, cat.UpdatedAt, cat.Gender, cat.VaccinationStatus,
	)
	return err
}

// UpdateCat memperbarui data kucing
func (r *catRepository) UpdateCat(cat *model.Cat) error {
	_, err := r.db.Exec(
		`UPDATE m_cat SET name = $1, breed = $2, age = $3, color = $4, description = $5,
		 adopted = $6, latitude = $7, longitude = $8, location_name = $9, photo_url = $10,
		 updated_at = $11, gender = $12, vaccination_status = $13 WHERE cat_id = $14`,
		cat.Name, cat.Breed, cat.Age, cat.Color, cat.Description, cat.Adopted,
		cat.Latitude, cat.Longitude, cat.LocationName, cat.PhotoURL, cat.UpdatedAt,
		cat.Gender, cat.VaccinationStatus, cat.CatID,
	)
	return err
}

// DeleteCat menghapus data kucing berdasarkan ID
func (r *catRepository) DeleteCat(catID string) error {
	_, err := r.db.Exec("DELETE FROM m_cat WHERE cat_id = $1", catID)
	return err
}

func NewCatRepository(db *sql.DB) catRepository {
	return &catRepository{
		db: db}
}
