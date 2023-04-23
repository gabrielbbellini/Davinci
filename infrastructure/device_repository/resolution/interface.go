package resolution

import (
	"davinci/domain/entities"
	"context"
)

type Repository interface {

	// GetAll return all resolution from the database.
	GetAll(ctx context.Context) ([]entities.Resolution, error)

	// GetById return a resolution by id.
	GetById(ctx context.Context, id int64) (entities.Resolution, error)
}
