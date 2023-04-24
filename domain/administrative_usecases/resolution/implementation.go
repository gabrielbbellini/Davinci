package resolution

import (
	"context"
	"davinci/domain/entities"
	resolution_mod "davinci/infrastructure/administrative_repository/resolution"
	"davinci/settings"
	"log"
)

type useCases struct {
	resolutionRepository resolution_mod.Repository
	settings             settings.Settings
}

func NewUseCases(settings settings.Settings, resolutionRepo resolution_mod.Repository) UseCases {
	return &useCases{
		resolutionRepository: resolutionRepo,
		settings:             settings,
	}
}

func (u useCases) GetAll(ctx context.Context) ([]entities.Resolution, error) {
	resolutions, err := u.resolutionRepository.GetAll(ctx)
	if err != nil {
		log.Println("[GetAll] Error GetAll")
		return nil, err
	}
	return resolutions, nil
}

func (u useCases) GetById(ctx context.Context, resolutionId int64) (*entities.Resolution, error) {
	resolution, err := u.resolutionRepository.GetById(ctx, resolutionId)
	if err != nil {
		log.Println("[GetAll] Error GetById")
		return nil, err
	}
	return resolution, nil
}
