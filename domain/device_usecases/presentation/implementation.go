package presentation

import (
	"base/domain/entities"
	"base/infrastructure/device_repository/presentation"
	"context"
)

type useCases struct {
	presentationRepo presentation.Repository
}

func NewUseCases(presentationRepo presentation.Repository) UseCases {
	return &useCases{
		presentationRepo: presentationRepo,
	}
}

func (u useCases) GetAll(ctx context.Context, idUser int64, idResolution int64) ([]entities.Presentation, error) {
	return u.presentationRepo.GetAll(ctx, idUser, idResolution)
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
