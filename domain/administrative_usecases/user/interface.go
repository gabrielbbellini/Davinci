package user

import "base/domain/entities"

type UseCases interface {
	Create(user entities.User) error
	Get(id int64) (*entities.User, error)
	GetAll() ([]entities.User, error)
	Update(user entities.User) error
	Delete(id int64) error
}
