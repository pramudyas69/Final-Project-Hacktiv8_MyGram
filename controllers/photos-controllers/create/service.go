package create

import "MyGram/models"

type Service interface {
	CreatePhotoService(input *InputCreatePhoto) (*models.Photo, string)
}

type service struct {
	repository Repository
}

func NewCreatePhotoService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreatePhotoService(input *InputCreatePhoto) (*models.Photo, string) {
	photos := models.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.PhotoUrl,
	}
	resultCreate, errCreate := s.repository.CreateRepositoryPhoto(&photos)

	return resultCreate, errCreate
}
