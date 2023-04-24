package presentation

import (
	"context"
	"davinci/domain/entities"
)

type Repository interface {
	// GetAll return all presentations from the database.
	GetAll(ctx context.Context, userId int64, idResolution int64) ([]entities.Presentation, error)

	// GetById return a presentation by id.
	GetById(ctx context.Context, id int64, userId int64) (*entities.Presentation, error)
}
