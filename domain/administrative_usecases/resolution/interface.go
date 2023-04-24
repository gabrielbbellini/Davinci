package resolution

import (
	"context"
	"davinci/domain/entities"
)

type UseCases interface {

	// GetAll return all resolution from the database.
	GetAll(ctx context.Context) ([]entities.Resolution, error)

	// GetById return a resolution by id.
	GetById(ctx context.Context, resolutionId int64) (*entities.Resolution, error)
}
