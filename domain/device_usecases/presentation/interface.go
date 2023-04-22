package presentation

import (
	"base/domain/entities"
	"context"
)

type UseCases interface {
	// GetAll return all devices.
	GetAll(ctx context.Context, idUser int64, idResolution int64) ([]entities.Presentation, error)

	// GetById return a device by id.
	GetById(ctx context.Context, id int64, idUser int64) (*entities.Presentation, error)
}
