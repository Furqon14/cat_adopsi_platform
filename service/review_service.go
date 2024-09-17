package service //

import (
	"cat_adoption_platform/model"
	"cat_adoption_platform/repository"
)

type ReviewService interface {
	Create(review model.Review) (model.Review, error)
	GetByID(id string) (model.Review, error)
	Update(review model.Review) (model.Review, error)
	Delete(id string) error
	GetAll() ([]model.Review, error)
}

type reviewService struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewService(reviewRepository repository.ReviewRepository) ReviewService {
	return &reviewService{reviewRepository: reviewRepository}
}

func (s *reviewService) Create(review model.Review) (model.Review, error) {
	return s.reviewRepository.Create(review)
}

func (s *reviewService) GetByID(id string) (model.Review, error) {
	return s.reviewRepository.GetByID(id)
}

func (s *reviewService) Update(review model.Review) (model.Review, error) {
	return s.reviewRepository.Update(review)
}

func (s *reviewService) Delete(id string) error {
	return s.reviewRepository.Delete(id)
}

func (s *reviewService) GetAll() ([]model.Review, error) {
	return s.reviewRepository.GetAll()
}
