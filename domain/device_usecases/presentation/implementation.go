package presentation

import (
	"context"
	"davinci/domain/entities"
	"davinci/infrastructure/device_repository/presentation"
)

type useCases struct {
	presentationRepo presentation.Repository
}

func NewUseCases(presentationRepo presentation.Repository) UseCases {
	return &useCases{
		presentationRepo: presentationRepo,
	}
}

func (u useCases) GetAll(ctx context.Context, userId int64, idResolution int64) ([]entities.Presentation, error) {
	return u.presentationRepo.GetAll(ctx, userId, idResolution)
}

func (u useCases) GetById(
	ctx context.Context,
	id int64,
	userId int64,
) (
	*entities.Presentation,
	error,
) {
	return u.presentationRepo.GetById(
		ctx,
		id,
		userId,
	)
}
