package presentation

import (
	"context"
	"davinci/domain/entities"
)

type UseCases interface {
	// GetAll return all devices.
	GetAll(ctx context.Context, userId int64, idResolution int64) ([]entities.Presentation, error)

	// GetById return a device by id.
	GetById(ctx context.Context, id int64, userId int64) (*entities.Presentation, error)
}
