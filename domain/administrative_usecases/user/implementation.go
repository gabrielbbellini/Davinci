package user

import (
	"base/domain/entities"
	"base/infrastructure/administrative_repository/user"
)

type useCase struct {
	repository user.Repository
}

func NewUseCase(repository user.Repository) UseCases {
	return &useCase{
		repository: repository,
	}
}

func (u useCase) Create(user entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (u useCase) Get(id int64) (*entities.User, error) {
	return u.repository.Get(id)
}

func (u useCase) GetAll() ([]entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u useCase) Update(user entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (u useCase) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
