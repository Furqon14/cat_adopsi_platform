package service

import (
	"cat_adoption_platform/model"
	"cat_adoption_platform/repository"
	"fmt"
)

type CatService interface {
	GetAllCats() ([]model.Cat, error)
	GetCatByID(id string) (*model.Cat, error)
	CreateCat(cat *model.Cat) (*model.Cat, error)
	DeleteCat(id string) error
	PostCatImages(payload []model.CatImage) ([]model.CatImage, error)
}

type catService struct {
	repo repository.CatRepository
}

func (s *catService) PostCatImages(payload []model.CatImage) ([]model.CatImage, error) {
	return s.repo.PostImages(payload)
}

// GetAllCats mengambil semua data kucing
func (s *catService) GetAllCats() ([]model.Cat, error) {
	cats, err := s.repo.GetAllCats()
	if err != nil {
		return nil, err
	}
	return cats, nil
}

// GetCatByID mengambil data kucing berdasarkan ID
func (s *catService) GetCatByID(id string) (*model.Cat, error) {
	return s.repo.GetCatByID(id)
}

// CreateCat menambahkan data kucing baru
func (s *catService) CreateCat(cat *model.Cat) (*model.Cat, error) {
	createdCat, err := s.repo.CreateCat(cat)
	if err != nil {
		// Tambahkan log untuk melacak error
		fmt.Println("Error creating cat in service:", err)
		return nil, err
	}
	return createdCat, nil
}

// DeleteCat menghapus data kucing berdasarkan ID
func (s *catService) DeleteCat(id string) error {
	return s.repo.DeleteCat(id)
}

func NewCatService(repo repository.CatRepository) CatService {
	return &catService{repo: repo}
}
