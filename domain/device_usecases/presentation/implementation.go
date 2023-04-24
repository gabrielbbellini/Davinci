package presentation

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	"davinci/infrastructure/device_repository/presentation"
	"davinci/view/http_error"
	"errors"
	"log"
)

type useCases struct {
	presentationRepo presentation.Repository
}

func NewUseCases(presentationRepo presentation.Repository) UseCases {
	return &useCases{
		presentationRepo: presentationRepo,
	}
}

func (u useCases) GetCurrentPresentation(ctx context.Context, deviceId int64) (*entities.Presentation, error) {
	currentPresentation, err := u.presentationRepo.GetCurrentPresentation(ctx, deviceId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetCurrentPresentation] Error GetCurrentPresentation", err)
		return nil, http_error.NewInternalServerError("Erro ao consultar apresentação atual.")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, http_error.NewBadRequestError("Não há nenhuma apresentação tocando no momento.")
	}

	return currentPresentation, nil
}
