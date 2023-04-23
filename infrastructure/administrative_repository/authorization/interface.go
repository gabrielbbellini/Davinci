package authorization

import (
	"davinci/domain/entities"
	"context"
)

type Repository interface {
	// Login check credentials on database and return error if not match.
	Login(ctx context.Context, credential entities.Credential) (*entities.User, error)
}
