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

func (u useCases) Create(ctx context.Context, presentation entities.Presentation, userId int64) (int64, error) {
	if presentation.Name = strings.TrimSpace(presentation.Name); presentation.Name == "" {
		log.Println("[Create] Error presentation.Name == \"\"")
		return 0, http_error.NewBadRequestError("Nome da apresentação deve ser informado.")
	}

	if presentation.Orientation != entities.OrientationLandscape && presentation.Orientation != entities.OrientationPortrait {
		log.Println("[Create] Error presentation.Name == \"\"")
		return 0, http_error.NewBadRequestError("Orientação informada não é válida.")
	}

	_, err := u.resolutionRepository.GetById(ctx, presentation.ResolutionId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[Create] Error GetById", err)
		return 0, http_error.NewInternalServerError("Erro ao consultar as resoluções disponíveis.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("[Create] Error resolution is not valid.", err)
		return 0, http_error.NewBadRequestError("Resolução não é válida.")
	}

	presentationId, err := u.presentationRepo.Create(ctx, presentation, userId)
	if err != nil {
		log.Println("[Create] Error Create", err)
		return 0, http_error.NewInternalServerError("Erro ao cadastrar apresentação.")
	}

	return presentationId, nil
}

func (u useCases) Update(ctx context.Context, presentationId int64, presentation entities.Presentation, userId int64) error {
	if presentation.Name = strings.TrimSpace(presentation.Name); presentation.Name == "" {
		log.Println("[Update] Error presentation.Name == \"\"")
		return http_error.NewBadRequestError("Nome da apresentação deve ser informado.")
	}

	if presentation.Orientation != entities.OrientationLandscape && presentation.Orientation != entities.OrientationPortrait {
		log.Println("[Update] Error presentation.Name == \"\"")
		return http_error.NewBadRequestError("Orientação informada não é válida.")
	}

	_, err := u.resolutionRepository.GetById(ctx, presentation.ResolutionId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[Update] Error GetById", err)
		return http_error.NewInternalServerError("Erro ao consultar as resoluções disponíveis.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("[Update] Error resolution is not valid.", err)
		return http_error.NewBadRequestError("Resolução não é válida.")
	}

	err = u.presentationRepo.Update(ctx, presentationId, presentation, userId)
	if err != nil {
		log.Println("[Update] Error Update", err)
		return http_error.NewInternalServerError("Erro ao editar apresentação.")
	}

	return nil
}

func (u useCases) Delete(ctx context.Context, presentationId int64, userId int64) error {
	_, err := u.presentationRepo.GetById(ctx, presentationId, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[Delete] Error GetById", err)
		return http_error.NewInternalServerError("Erro ao deletar apresentação.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return http_error.NewBadRequestError("Apresentação não encontrada.")
	}

	err = u.presentationRepo.Delete(ctx, presentationId, userId)
	if err != nil {
		log.Println("[Delete] Error Delete", err)
		return http_error.NewInternalServerError("Erro ao deletar dispositivo.")
	}

	return nil
}

func (u useCases) GetAll(ctx context.Context, userId int64) ([]entities.Presentation, error) {
	devices, err := u.presentationRepo.GetAll(ctx, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetAll] Error GetAll", err)
		return nil, http_error.NewInternalServerError("Erro ao consultar as apresentações.")
	}

	return devices, nil
}

func (u useCases) GetById(ctx context.Context, id int64, userId int64) (*entities.Presentation, error) {
	presentation, err := u.presentationRepo.GetById(ctx, id, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetById] Error GetById", err)
		return nil, http_error.NewInternalServerError("Erro ao consultar apresentação.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, http_error.NewBadRequestError("Apresentação não encontrada.")
	}

	return presentation, nil
}
