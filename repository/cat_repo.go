package repository

import (
	"cat_adoption_platform/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type CatRepository interface {
	GetAllCats() ([]model.Cat, error)
	GetCatByID(catID string) (*model.Cat, error)
	CreateCat(cat *model.Cat) (*model.Cat, error)
	DeleteCat(catID string) error
	PostImages(payload []model.CatImage) ([]model.CatImage, error)
}
type catRepository struct {
	db *sql.DB
}

func (r *catRepository) PostImages(payload []model.CatImage) ([]model.CatImage, error) {
	// query := `INSERT INTO t_cat_image (id, cat_id, image_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, cat_id, image_url, created_at, updated_at`
	// err := r.db.QueryRow(query, payload.ID, payload.CatID, payload.URL).Scan(&payload.ID, &payload.CatID, &payload.URL, &payload.CreatedAt, &payload.UpdatedAt)
	// if err != nil && err != sql.ErrNoRows {
	// 	log.Fatal(err)
	// }
	return payload, nil
}

// GetAllCats mengambil semua data kucing
func (r *catRepository) GetAllCats() ([]model.Cat, error) {
	query := `SELECT cat_id, name, breed, age, color, description, adopted, latitude, longitude, location_name, photo_url, gender, vaccination_status, created_at, updated_at FROM m_cat`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []model.Cat
	for rows.Next() {
		var cat model.Cat
		if err := rows.Scan(&cat.CatID, &cat.Name, &cat.Breed, &cat.Age, &cat.Color, &cat.Description, &cat.Adopted, &cat.Latitude, &cat.Longitude, &cat.LocationName, &cat.PhotoURL, &cat.Gender, &cat.VaccinationStatus, &cat.CreatedAt, &cat.UpdatedAt); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
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
func (r *catRepository) CreateCat(cat *model.Cat) (*model.Cat, error) {
	query := `
        INSERT INTO m_cat (name, breed, age, color, description, adopted, latitude, longitude, location_name, photo_url, gender, vaccination_status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
        RETURNING cat_id, created_at, updated_at`

	row := r.db.QueryRow(query, cat.Name, cat.Breed, cat.Age, cat.Color, cat.Description, cat.Adopted, cat.Latitude, cat.Longitude, cat.LocationName, cat.PhotoURL, cat.Gender, cat.VaccinationStatus, time.Now(), time.Now())

	if err := row.Scan(&cat.CatID, &cat.CreatedAt, &cat.UpdatedAt); err != nil {
		// Tambahkan log untuk melacak error
		fmt.Println("Error inserting cat in repository:", err)
		return nil, err
	}

	return cat, nil
}

// DeleteCat menghapus data kucing berdasarkan ID
func (r *catRepository) DeleteCat(catID string) error {
	query := "DELETE FROM m_cat WHERE cat_id = $1"
	result, err := r.db.Exec(query, catID)
	if err != nil {
		log.Println("Error deleting cat in repository:", err)
		return err
	}

	// Periksa apakah ada row yang terhapus
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected, cat not found")
	}

	return nil
}

func NewCatRepository(db *sql.DB) CatRepository {
	return &catRepository{db: db}
}
