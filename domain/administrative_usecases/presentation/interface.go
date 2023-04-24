package presentation

import (
	"context"
	"davinci/domain/entities"
)

type UseCases interface {
	// Create new device.
	Create(ctx context.Context, presentation entities.Presentation, userId int64) error

	// Update update a device.
	Update(ctx context.Context, presentation entities.Presentation, userId int64) error

	// Delete remove a device.
	Delete(ctx context.Context, presentation entities.Presentation, userId int64) error

	// GetAll return all devices.
	GetAll(ctx context.Context, userId int64) ([]entities.Presentation, error)

	// GetById return a device by id.
	GetById(ctx context.Context, id int64, userId int64) (*entities.Presentation, error)
}
