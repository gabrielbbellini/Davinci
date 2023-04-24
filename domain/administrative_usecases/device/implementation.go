package device

import (
	"database/sql"
	"davinci/domain/entities"
	deviceRepository "davinci/infrastructure/administrative_repository/device"
	"davinci/infrastructure/administrative_repository/resolution"
	"davinci/settings"
	"errors"

	"context"
	"davinci/view/http_error"
	"log"
	"strings"
)

type useCases struct {
	deviceRepository     deviceRepository.Repository
	resolutionRepository resolution.Repository
	settings             settings.Settings
}

func NewUseCases(settings settings.Settings, deviceRepository deviceRepository.Repository, repositoryRepository resolution.Repository) UseCases {
	return &useCases{
		deviceRepository:     deviceRepository,
		resolutionRepository: repositoryRepository,
		settings:             settings,
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
		log.Println("[Create] Error GetAll", err)
		return http_error.NewInternalServerError("Erro ao consultar as resoulões disponíveis.")
	}

	var isResolutionValid bool
	for _, it := range resolutions {
		if it.Id == device.ResolutionId {
			isResolutionValid = true
			break
		}
	}

	if !isResolutionValid {
		log.Println("[Create] Error !isResolutionValid", err)
		return http_error.NewBadRequestError("Resolução não é válida.")
	}

	_, err = u.deviceRepository.GetDeviceByName(ctx, device.Name, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[Create] Error GetDeviceByName")
		return http_error.NewInternalServerError("Erro ao cadastrar dispositivo")
	}
	if err == nil {
		return http_error.NewBadRequestError("Já existe um dispositivo com o mesmo nome.")
	}

	err = u.deviceRepository.Create(ctx, device, userId)
	if err != nil {
		log.Println("[Create] Error Create")
		return http_error.NewInternalServerError("Erro ao cadastrar dispositivo")
	}

	return nil
}

func (u useCases) Update(ctx context.Context, deviceId int64, device entities.Device, userId int64) error {
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
		if it.Id == device.ResolutionId {
			isResolutionValid = true
			break
		}
	}

	if !isResolutionValid {
		log.Println("[Update] error !isResolutionValid", err)
		return http_error.NewBadRequestError("Resolução não é válida.")
	}

	deviceOld, err := u.deviceRepository.GetById(ctx, deviceId, userId)
	if err != nil {
		log.Println("[Update] Error GetById", err)
		return http_error.NewBadRequestError("Dispositivo não encontrado")
	}

	deviceByName, err := u.deviceRepository.GetDeviceByName(ctx, device.Name, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[Update] Error GetDeviceByName")
		return http_error.NewInternalServerError("Erro ao cadastrar dispositivo")
	}
	if err == nil && deviceByName.Id != deviceOld.Id {
		return http_error.NewBadRequestError("Já existe um dispositivo com o mesmo nome.")
	}

	err = u.deviceRepository.Update(ctx, deviceId, device, userId)
	if err != nil {
		log.Println("[Update] Error GetById", err)
		return http_error.NewInternalServerError("Erro ao editar dispositivo")
	}

	return nil
}

func (u useCases) Delete(ctx context.Context, deviceId int64, userId int64) error {
	_, err := u.deviceRepository.GetById(ctx, deviceId, userId)
	if err != nil {
		log.Println("[Delete] Error GetById", err)
		return http_error.NewBadRequestError("Dispositivo não encontrado.")
	}

	if err = u.deviceRepository.Delete(ctx, deviceId, userId); err != nil {
		log.Println("[Delete] Error Delete", err)
		return http_error.NewInternalServerError("Erro ao deletar dispositivo.")
	}

	return nil
}

func (u useCases) GetAll(ctx context.Context, userId int64) ([]entities.Device, error) {
	devices, err := u.deviceRepository.GetAll(ctx, userId)
	if err != nil {
		log.Println("[GetAll] Error GetAll", err)
		return nil, http_error.NewInternalServerError("Erro ao listar dispositivos.")
	}

	return devices, nil
}

func (u useCases) GetById(ctx context.Context, deviceId int64, userId int64) (*entities.Device, error) {
	device, err := u.deviceRepository.GetById(ctx, deviceId, userId)
	if err != nil {
		log.Println("[GetById] Error GetById", err)
		return nil, http_error.NewBadRequestError("Dispositivo não encontrado.")
	}

	return device, nil
}
