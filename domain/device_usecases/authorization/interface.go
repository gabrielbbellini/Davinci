package authorization

import (
	"context"
	"davinci/domain/entities"
)

type UseCases interface {
	// Login will check credential and return error if credential is not valid.
	Login(ctx context.Context, credential entities.Credential) (*entities.Device, error)
}
