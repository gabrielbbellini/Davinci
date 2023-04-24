package presentation

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	presentation_mod "davinci/infrastructure/administrative_repository/presentation"
	"davinci/infrastructure/administrative_repository/resolution"
	"davinci/settings"
	"davinci/view/http_error"
	"errors"
	"log"
	"strings"
)

type useCases struct {
	presentationRepo     presentation_mod.Repository
	resolutionRepository resolution.Repository
	settings             settings.Settings
}

func NewUseCases(settings settings.Settings, resolutionRepository resolution.Repository, presentationRepo presentation_mod.Repository) UseCases {
	return &useCases{
		presentationRepo:     presentationRepo,
		resolutionRepository: resolutionRepository,
		settings:             settings,
	}
}

func (u useCases) Create(ctx context.Context, presentation entities.Presentation, userId int64) error {
	if presentation.Name = strings.TrimSpace(presentation.Name); presentation.Name == "" {
		log.Println("[Create] Error presentation.Name == \"\"")
		return http_error.NewBadRequestError("Nome da apresentação deve ser informado.")
	}

	if presentation.Orientation != entities.OrientationLandscape && presentation.Orientation != entities.OrientationPortrait {
		log.Println("[Create] Error presentation.Name == \"\"")
		return http_error.NewBadRequestError("Orientação informada não é válida.")
	}

	resolutions, err := u.resolutionRepository.GetAll(ctx)
	if err != nil {
		log.Println("[Create] Error GetAll", err)
		return http_error.NewInternalServerError("Erro ao consultar as resoluções disponíveis.")
	}

	var isResolutionValid bool
	for _, it := range resolutions {
		if it.Id == presentation.ResolutionId {
			isResolutionValid = true
			break
		}
	}

	if !isResolutionValid {
		log.Println("[Create] Error !isResolutionValid", err)
		return http_error.NewBadRequestError("Resolução não é válida.")
	}

	err = u.presentationRepo.Create(ctx, presentation, userId)
	if err != nil {
		log.Println("[Create] Error Create", err)
		return http_error.NewInternalServerError("Erro ao cadastrar apresentaçao.")
	}

	return nil
}

func (u useCases) Update(ctx context.Context, presentation entities.Presentation, userId int64) error {
	return u.presentationRepo.Update(ctx, presentation, userId)
}

func (u useCases) Delete(ctx context.Context, presentation entities.Presentation, userId int64) error {
	return u.presentationRepo.Delete(ctx, presentation, userId)
}

func (u useCases) GetAll(ctx context.Context, userId int64) ([]entities.Presentation, error) {
	devices, err := u.presentationRepo.GetAll(ctx, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetAll] Error GetAll", err)
		return nil, http_error.NewInternalServerError("Erro ao consultar os dispositivos.")
	}

	return devices, nil
}

func (u useCases) GetById(ctx context.Context, id int64, userId int64) (*entities.Presentation, error) {
	presentation, err := u.presentationRepo.GetById(ctx, id, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetById] Error GetById", err)
		return nil, http_error.NewInternalServerError("Erro ao consultar dispositivo.")
	}

	return presentation, nil
}
