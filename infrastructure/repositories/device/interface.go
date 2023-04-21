package device

import (
	"base/domain/entities"
	"context"
)

type Repository interface {
	// Create insert a new device in the database.
	Create(ctx context.Context, device entities.Device) error

	// Update a device in the database.
	Update(ctx context.Context, device entities.Device) error

	// Delete remove a device from the database.
	Delete(ctx context.Context, device entities.Device) error

	// GetAll return all devices from the database.
	GetAll(ctx context.Context) ([]entities.Device, error)

	// GetById return a device by id.
	GetById(ctx context.Context, id int64) (entities.Device, error)
}
