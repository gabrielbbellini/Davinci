package authorization

import (
	"context"
	"davinci/domain/entities"
	"davinci/infrastructure/administrative_repository/authorization"
	"davinci/settings"
	"davinci/view/http_error"
	"strings"
)

type useCases struct {
	repository authorization.Repository
	settings   settings.Settings
}

func NewUseCases(settings settings.Settings, repository authorization.Repository) UseCases {
	return &useCases{
		repository: repository,
		settings:   settings,
	}
}

func (u useCases) Login(ctx context.Context, credential entities.Credential) (*entities.User, error) {
	if credential.Email = strings.TrimSpace(credential.Email); credential.Email == "" {
		return nil, http_error.NewForbiddenError(http_error.ForbiddenMessage)
	}

	if credential.Password = strings.TrimSpace(credential.Password); credential.Password == "" {
		return nil, http_error.NewForbiddenError(http_error.ForbiddenMessage)
	}

	return u.repository.Login(ctx, credential)
}
