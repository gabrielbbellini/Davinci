package presentation

import (
	"context"
	"davinci/domain/entities"
	"davinci/infrastructure/administrative_repository/presentation"
	"davinci/settings"
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
	return u.presentationRepo.GetAll(ctx, idUser)
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
