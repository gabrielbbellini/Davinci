package resolution

import (
	"base/domain/entities"
	"base/infrastructure/repositories/resolution"
	"context"
)

type useCases struct {
	resolutionRepo resolution.Repository
}

func NewUseCases(resolutionRepo resolution.Repository) UseCases {
	return &useCases{
		resolutionRepo: resolutionRepo,
	}
}

func (u useCases) GetAll(ctx context.Context) ([]entities.Resolution, error) {
	return u.resolutionRepo.GetAll(ctx)
}

func (u useCases) GetById(ctx context.Context, id int64) (entities.Resolution, error) {
	return u.resolutionRepo.GetById(ctx, id)
}
