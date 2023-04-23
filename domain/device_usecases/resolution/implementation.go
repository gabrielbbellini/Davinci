package resolution

import (
	"davinci/domain/entities"
	"davinci/infrastructure/device_repository/resolution"
	"context"
)

type useCases struct {
	resolutionRepository resolution.Repository
}

func NewUseCases(resolutionRepository resolution.Repository) UseCases {
	return &useCases{
		resolutionRepository: resolutionRepository,
	}
}

func (u useCases) GetAll(ctx context.Context) ([]entities.Resolution, error) {
	return u.resolutionRepository.GetAll(ctx)
}

func (u useCases) GetById(ctx context.Context, id int64) (entities.Resolution, error) {
	return u.resolutionRepository.GetById(ctx, id)
}
