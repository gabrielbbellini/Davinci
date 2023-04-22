package presentation

import (
	"base/domain/entities"
	"context"
)

type Repository interface {
	// GetAll return all presentations from the database.
	GetAll(ctx context.Context, idUser int64, idResolution int64) ([]entities.Presentation, error)

	// GetById return a presentation by id.
	GetById(ctx context.Context, id int64, idUser int64) (*entities.Presentation, error)
}
