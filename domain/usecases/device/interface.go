package device

import (
	"base/domain/entities"
	"context"
)

type UseCases interface {
	// Create new device.
	Create(ctx context.Context, device entities.Device, idUser int64) error
	// Update update a device.
	Update(ctx context.Context, device entities.Device, idUser int64) error
	// Delete remove a device.
	Delete(ctx context.Context, device entities.Device, idUser int64) error

	// GetAll return all devices.
	GetAll(ctx context.Context, idUser int64) ([]entities.Device, error)

	// GetById return a device by id.
	GetById(ctx context.Context, id int64, idUser int64) (entities.Device, error)
}
