package authorization

import (
	"davinci/domain/entities"
	"context"
)

type UseCases interface {
	// Login will check credential and return error if credential is not valid.
	Login(ctx context.Context, credential entities.Credential) (*entities.User, error)
}
