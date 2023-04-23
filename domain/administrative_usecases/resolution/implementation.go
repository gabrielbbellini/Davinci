package resolution

import (
	"context"
	"davinci/domain/entities"
	"davinci/infrastructure/device_repository/resolution"
	"davinci/settings"
)

type useCases struct {
	resolutionRepository resolution.Repository
	settings             settings.Settings
}

func NewUseCases(settings settings.Settings, resolutionRepo resolution.Repository) UseCases {
	return &useCases{
		resolutionRepository: resolutionRepo,
		settings:             settings,
	}
}

func (u useCases) GetAll(ctx context.Context) ([]entities.Resolution, error) {
	return u.resolutionRepository.GetAll(ctx)
}

func (u useCases) GetById(ctx context.Context, id int64) (entities.Resolution, error) {
	return u.resolutionRepository.GetById(ctx, id)
}
