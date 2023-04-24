package presentation

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	"davinci/infrastructure/administrative_repository/presentation"
	"davinci/settings"
	"davinci/view/http_error"
	"errors"
	"log"
)

type useCases struct {
	presentationRepo presentation.Repository
	settings         settings.Settings
}

func NewUseCases(settings settings.Settings, presentationRepo presentation.Repository) UseCases {
	return &useCases{
		presentationRepo: presentationRepo,
		settings:         settings,
	}
}

func (u useCases) Create(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	return u.presentationRepo.Create(ctx, presentation, idUser)
}

func (u useCases) Update(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	return u.presentationRepo.Update(ctx, presentation, idUser)
}

func (u useCases) Delete(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	return u.presentationRepo.Delete(ctx, presentation, idUser)
}

func (u useCases) GetAll(ctx context.Context, idUser int64) ([]entities.Presentation, error) {
	devices, err := u.presentationRepo.GetAll(ctx, idUser)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetAll] Error GetAll", err)
		return nil, http_error.NewInternalServerError("Erro ao consultar os dispositivos.")
	}

	return devices, nil
}

func (u useCases) GetById(
	ctx context.Context,
	id int64,
	idUser int64,
) (
	*entities.Presentation,
	error,
) {
	return u.presentationRepo.GetById(
		ctx,
		id,
		idUser,
	)
}
