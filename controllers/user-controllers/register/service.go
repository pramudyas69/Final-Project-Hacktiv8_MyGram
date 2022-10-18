package register

import "MyGram/models"

type Service interface {
	RegisterService(input *InputRegister) (*models.User, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*models.User, string) {
	users := models.User{
		Age:      input.Age,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&users)

	return resultRegister, errRegister
}
