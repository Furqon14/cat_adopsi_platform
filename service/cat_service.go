package service

import (
	"cat_adoption_platform/model"
	"cat_adoption_platform/repository"
)

type CatService interface {
	GetAllCats() ([]model.Cat, error)
	GetCatByID(id string) (*model.Cat, error)
	CreateCat(cat *model.Cat) error
	UpdateCat(cat *model.Cat) error
	DeleteCat(id string) error
}

type catService struct {
	repo repository.CatRepository
}

// GetAllCats mengambil semua data kucing
func (s *catService) GetAllCats() ([]model.Cat, error) {
	return s.repo.GetAllCats()
}

// GetCatByID mengambil data kucing berdasarkan ID
func (s *catService) GetCatByID(id string) (*model.Cat, error) {
	return s.repo.GetCatByID(id)
}

// CreateCat menambahkan data kucing baru
func (s *catService) CreateCat(cat *model.Cat) error {
	return s.repo.CreateCat(cat)
}

// UpdateCat memperbarui data kucing
func (s *catService) UpdateCat(cat *model.Cat) error {
	return s.repo.UpdateCat(cat)
}

// DeleteCat menghapus data kucing berdasarkan ID
func (s *catService) DeleteCat(id string) error {
	return s.repo.DeleteCat(id)
}

func NewcatService(repo *repository.CatRepository) *catService {
	return &catService{repo: repo}
}
