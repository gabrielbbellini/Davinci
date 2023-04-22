package device

import (
	"base/domain/entities"
	deviceRepository "base/infrastructure/administrative_repository/device"
	"base/infrastructure/administrative_repository/resolution"

	"base/view/http_error"
	"context"
	"log"
	"strings"
)

type useCases struct {
	deviceRepository     deviceRepository.Repository
	resolutionRepository resolution.Repository
}

func NewUseCases(deviceRepository deviceRepository.Repository, repositoryRepository resolution.Repository) UseCases {
	return &useCases{
		deviceRepository:     deviceRepository,
		resolutionRepository: repositoryRepository,
	}
}

func (u useCases) Create(ctx context.Context, device entities.Device, userId int64) error {
	if device.Name = strings.TrimSpace(device.Name); device.Name == "" {
		log.Println("[Create] Error device.Name == \"\"")
		return http_error.NewBadRequestError("Nome do dispositivo deve ser informado.")
	}

	if device.Orientation != entities.OrientationLandscape && device.Orientation != entities.OrientationPortrait {
		log.Println("[Create] Error device.Name == \"\"")
		return http_error.NewBadRequestError("Orientação informada não é válida.")
	}

	resolutions, err := u.resolutionRepository.GetAll(ctx)
	if err != nil {
		log.Println("[Create] error GetAll", err)
		return http_error.NewInternalServerError("Erro ao consultar as resoulões disponíveis.")
	}

	var isResolutionValid bool
	for _, it := range resolutions {
		if it == device.Resolution {
			isResolutionValid = true
			break
		}
	}

	if !isResolutionValid {
		log.Println("[Create] error !isResolutionValid", err)
		return http_error.NewInternalServerError("Resolução não é válida.")
	}

	return u.deviceRepository.Create(ctx, device, userId)
}

func (u useCases) Update(ctx context.Context, device entities.Device, userId int64) error {
	if device.Name = strings.TrimSpace(device.Name); device.Name == "" {
		log.Println("[Update] Error device.Name == \"\"")
		return http_error.NewBadRequestError("Nome do dispositivo deve ser informado.")
	}

	if device.Orientation != entities.OrientationLandscape && device.Orientation != entities.OrientationPortrait {
		log.Println("[Update] Error device.Name == \"\"")
		return http_error.NewBadRequestError("Orientação informada não é válida.")
	}

	resolutions, err := u.resolutionRepository.GetAll(ctx)
	if err != nil {
		log.Println("[Update] error GetAll", err)
		return http_error.NewInternalServerError("Erro ao consultar as resoulões disponíveis.")
	}

	var isResolutionValid bool
	for _, it := range resolutions {
		if it == device.Resolution {
			isResolutionValid = true
			break
		}
	}

	if !isResolutionValid {
		log.Println("[Update] error !isResolutionValid", err)
		return http_error.NewInternalServerError("Resolução não é válida.")
	}

	_, err = u.deviceRepository.GetById(ctx, device.Id, userId)
	if err != nil {
		log.Println("[Update] Error GetById", err)
		return http_error.NewInternalServerError("Device não encontrado")
	}

	return u.deviceRepository.Update(ctx, device, userId)
}

func (u useCases) Delete(ctx context.Context, deviceId int64, userId int64) error {
	_, err := u.deviceRepository.GetById(ctx, deviceId, userId)
	if err != nil {
		log.Println("[Delete] Error GetById", err)
		return http_error.NewBadRequestError("Dispositivo não encontrado")
	}

	return u.deviceRepository.Delete(ctx, deviceId, userId)
}

func (u useCases) GetAll(ctx context.Context, userId int64) ([]entities.Device, error) {
	return u.deviceRepository.GetAll(ctx, userId)
}

func (u useCases) GetById(ctx context.Context, deviceId int64, userId int64) (*entities.Device, error) {
	device, err := u.deviceRepository.GetById(ctx, deviceId, userId)
	if err != nil {
		log.Println("[GetById] Error GetById", err)
		return nil, http_error.NewBadRequestError("Dispositivo não encontrado")
	}

	return device, nil
}
