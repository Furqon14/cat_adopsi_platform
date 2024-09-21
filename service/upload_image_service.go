package service

import (
	"cat_adoption_platform/model"
	"cat_adoption_platform/repository"
	"cat_adoption_platform/uploaders"
	"mime/multipart"
)

type UploadImageService interface {
	Upload(file multipart.File, filename string) (model.Image, error)
}

type uploadImageService struct {
	repo     repository.UploadImageRepository
	uploader uploaders.LocalImageUploader
}

func (us *uploadImageService) Upload(file multipart.File, filename string) (model.Image, error) {
	url, err := us.uploader.Upload(file, filename)
	if err != nil {
		return model.Image{}, err
	}

	image := model.Image{
		Filename: filename,
		URL:      url,
	}

	err = us.repo.Save(image)
	if err != nil {
		return model.Image{}, err
	}

	return image, nil
}

func NewUploadImageService(
	repo repository.UploadImageRepository,
	uploader uploaders.LocalImageUploader,
) UploadImageService {
	return &uploadImageService{
		repo:     repo,
		uploader: uploader,
	}
}
