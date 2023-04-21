package user

import (
	"base/domain/entities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, user entities.User) error
	Get(ctx context.Context, id int64) (*entities.User, error)
	GetAll(ctx context.Context) ([]entities.User, error)
	Update(ctx context.Context, user entities.User) error
	Delete(ctx context.Context, id int64) error
}
