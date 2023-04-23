package user

import (
	"context"
	"davinci/domain/entities"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}
