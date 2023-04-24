package resolution

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	resolution_mod "davinci/infrastructure/administrative_repository/resolution"
	"errors"
	"log"
)

type useCases struct {
	resolutionRepository resolution_mod.Repository
}

func NewUseCases(resolutionRepository resolution_mod.Repository) UseCases {
	return &useCases{
		resolutionRepository: resolutionRepository,
	}
}

func (u useCases) GetAll(ctx context.Context) ([]entities.Resolution, error) {
	resolutions, err := u.resolutionRepository.GetAll(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetAll] Error GetAll", err)
		return nil, err
	}

	return resolutions, nil
}

func (u useCases) GetById(ctx context.Context, resolutionId int64) (*entities.Resolution, error) {
	resolution, err := u.resolutionRepository.GetById(ctx, resolutionId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("[GetById] Error GetById", err)
		return nil, err
	}

	return &resolution, nil
}
