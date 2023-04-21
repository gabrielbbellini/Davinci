package device

import (
	"base/domain/entities"
	"context"
)

type UseCases interface {
	// Create new device.
	Create(ctx context.Context, device entities.Device) error

	// Update update a device.
	Update(ctx context.Context, device entities.Device) error

	// Delete remove a device.
	Delete(ctx context.Context, device entities.Device) error

	// GetAll return all devices.
	GetAll(ctx context.Context) ([]entities.Device, error)

	// GetById return a device by id.
	GetById(ctx context.Context, id int64) (entities.Device, error)
}
