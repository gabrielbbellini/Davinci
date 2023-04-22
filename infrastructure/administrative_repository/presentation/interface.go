package presentation

import (
	"base/domain/entities"
	"context"
)

type Repository interface {
	// Create insert a new presentation in the database.
	Create(ctx context.Context, presentation entities.Presentation, idUser int64) error

	// Update a presentation in the database.
	Update(ctx context.Context, presentation entities.Presentation, idUser int64) error

	// Delete remove a presentation from the database.
	Delete(ctx context.Context, presentation entities.Presentation, idUser int64) error

	// GetAll return all presentations from the database.
	GetAll(ctx context.Context, idUser int64) ([]entities.Presentation, error)

	// GetById return a presentation by id.
	GetById(ctx context.Context, id int64, idUser int64) (*entities.Presentation, error)
}
