package device

import (
	"base/domain/entities"
	"base/infrastructure/repositories/device"
	"context"
)

type useCases struct {
	deviceRepo device.Repository
}

func NewUseCases(deviceRepo device.Repository) UseCases {
	return &useCases{
		deviceRepo: deviceRepo,
	}
}

func (u useCases) Create(ctx context.Context, device entities.Device) error {
	//TODO implement me
	panic("implement me")
}

func (u useCases) Update(ctx context.Context, device entities.Device) error {
	//TODO implement me
	panic("implement me")
}

func (u useCases) Delete(ctx context.Context, device entities.Device) error {
	//TODO implement me
	panic("implement me")
}

func (u useCases) GetAll(ctx context.Context) ([]entities.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (u useCases) GetById(ctx context.Context, id int) (entities.Device, error) {
	//TODO implement me
	panic("implement me")
}
