package user

import "base/domain/entities"

type Repository interface {
	GetByEmail(email string) (*entities.User, error)
}
