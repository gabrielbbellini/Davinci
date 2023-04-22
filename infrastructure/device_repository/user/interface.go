package user

import (
	"base/domain/entities"
	"context"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}
