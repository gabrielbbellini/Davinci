package device_presentation

import (
	"context"
	"database/sql"
	"davinci/infrastructure/administrative_repository/device"
	"davinci/infrastructure/administrative_repository/device_presentation"
	"davinci/infrastructure/administrative_repository/presentation"
	"davinci/settings"
	"davinci/view/http_error"
	"errors"
	"log"
)

type useCases struct {
	repository             device_presentation.Repository
	deviceRepository       device.Repository
	presentationRepository presentation.Repository
	settings               settings.Settings
}

func NewUseCases(
	settings settings.Settings,
	repository device_presentation.Repository,
	deviceRepository device.Repository,
	presentationRepository presentation.Repository,
) UseCases {
	return &useCases{
		repository:             repository,
		deviceRepository:       deviceRepository,
		presentationRepository: presentationRepository,
		settings:               settings,
	}
}

func (u useCases) Relate(ctx context.Context, userId int64, deviceId int64, presentationId int64) error {
	_, err := u.deviceRepository.GetById(ctx, deviceId, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[Relate] Error GetById", err)
		return http_error.NewInternalServerError("Erro ao consultar dispositivo.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return http_error.NewBadRequestError("Dispositivo não encontrado.")
	}

	_, err = u.presentationRepository.GetById(ctx, presentationId, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[Relate] Error GetById", err)
		return http_error.NewInternalServerError("Erro ao consultar apresentação.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return http_error.NewBadRequestError("Apresentação não encontrada.")
	}

	err = u.repository.Relate(ctx, deviceId, presentationId)
	if err != nil {
		log.Println("[Relate] Error Relate", err)
		return http_error.NewInternalServerError("Erro ao relacionar apresentação ao dispositivo.")
	}

	return nil
}

func (u useCases) SetCurrentPresentation(ctx context.Context, userId int64, deviceId int64, presentationId int64) error {
	_, err := u.deviceRepository.GetById(ctx, deviceId, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[SetCurrentPresentation] Error GetById", err)
		return http_error.NewInternalServerError("Erro ao consultar dispositivo.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return http_error.NewBadRequestError("Dispositivo não encontrado.")
	}

	_, err = u.presentationRepository.GetById(ctx, presentationId, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[SetCurrentPresentation] Error GetById", err)
		return http_error.NewInternalServerError("Erro ao consultar apresentação.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return http_error.NewBadRequestError("Apresentação não encontrada.")
	}

	err = u.repository.SetCurrentPresentation(ctx, deviceId, presentationId)
	if err != nil {
		log.Println("[SetCurrentPresentation] Error SetCurrentPresentation", err)
		return http_error.NewInternalServerError("Erro ao alterar a apresentação atual do dispositivo.")
	}

	return nil
}

func (u useCases) GetCurrentPresentation(ctx context.Context, userId int64, deviceId int64) (int64, error) {
	_, err := u.deviceRepository.GetById(ctx, deviceId, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetCurrentPresentation] Error GetById", err)
		return 0, http_error.NewInternalServerError("Erro ao consultar dispositivo.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return 0, http_error.NewBadRequestError("Dispositivo não encontrado.")
	}

	presentationId, err := u.repository.GetCurrentPresentation(ctx, deviceId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetCurrentPresentation] Error GetCurrentPresentation", err)
		return 0, http_error.NewInternalServerError("Erro ao consultar apresentação atual.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return 0, http_error.NewBadRequestError("Não há nenhuma apresentação tocando no momento.")
	}

	return presentationId, nil
}
