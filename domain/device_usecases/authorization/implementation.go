package authorization

import (
	"base/domain/entities"
	"golang.org/x/crypto/bcrypt"

	"base/infrastructure/device_repository/authorization"
	"base/infrastructure/device_repository/device"
	user_repository "base/infrastructure/device_repository/user"
	"base/view/http_error"
	"context"
	"log"
	"strings"
)

type useCases struct {
	repository       authorization.Repository
	userRepository   user_repository.Repository
	deviceRepository device.Repository
}

func NewUseCases(
	repository authorization.Repository,
	userRepository user_repository.Repository,
	deviceRepository device.Repository,
) UseCases {
	return &useCases{
		repository:       repository,
		userRepository:   userRepository,
		deviceRepository: deviceRepository,
	}
}

func (u useCases) Login(ctx context.Context, credential entities.Credential) (*entities.Device, error) {
	if strings.TrimSpace(credential.Email) == "" {
		return nil, http_error.NewForbiddenError("Credenciais inválidas.")
	}

	if strings.TrimSpace(credential.Password) == "" {
		return nil, http_error.NewForbiddenError("Credenciais inválidas.")
	}

	user, err := u.userRepository.GetByEmail(credential.Email)
	if err != nil {
		log.Println("[Login] Error GetByEmail", err)
		return nil, http_error.NewForbiddenError("Credenciais inválidas.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Credential.Password), []byte(credential.Password))
	if err != nil {
		log.Println("[Login] Error CompareHashAndPassword", err)
		return nil, http_error.NewForbiddenError("Credenciais inválidas.")
	}

	return nil, err
}
