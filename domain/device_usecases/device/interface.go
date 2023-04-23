package device

import (
	"context"
	"davinci/domain/entities"
)

type UseCases interface {
	// Create new device.
	Create(ctx context.Context, device entities.Device, userId int64) error

	// Update update a device.
	Update(ctx context.Context, device entities.Device, userId int64) error

	// Delete remove a device.
	Delete(ctx context.Context, device entities.Device, userId int64) error

	// GetAll return all devices.
	GetAll(ctx context.Context, userId int64) ([]entities.Device, error)

	// GetById return a device by id.
	GetById(ctx context.Context, deviceId int64, userId int64) (entities.Device, error)
}
