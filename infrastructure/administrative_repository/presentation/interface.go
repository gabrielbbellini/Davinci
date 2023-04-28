package presentation

import (
	"context"
	"davinci/domain/entities"
)

type Repository interface {
	// Create insert a new presentation in the database.
	Create(ctx context.Context, presentation entities.Presentation, userId int64) (int64, error)

	// Update a presentation in the database.
	Update(ctx context.Context, presentationId int64, presentation entities.Presentation, userId int64) error

	// Delete remove a presentation from the database.
	Delete(ctx context.Context, presentationId int64, userId int64) error

	// GetAll return all presentations from the database.
	GetAll(ctx context.Context, userId int64) ([]entities.Presentation, error)

	// GetById return a presentation by id.
	GetById(ctx context.Context, presentationId int64, userId int64) (*entities.Presentation, error)
}
