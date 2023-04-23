package device

import (
	"context"
	"davinci/domain/entities"
	"davinci/infrastructure/device_repository/device"
)

type useCases struct {
	deviceRepo device.Repository
}

func (u useCases) Create(ctx context.Context, device entities.Device, userId int64) error {
	//TODO implement me
	panic("implement me")
}

func (u useCases) Update(ctx context.Context, device entities.Device, userId int64) error {
	//TODO implement me
	panic("implement me")
}

func (u useCases) Delete(ctx context.Context, device entities.Device, userId int64) error {
	//TODO implement me
	panic("implement me")
}

func (u useCases) GetAll(ctx context.Context, userId int64) ([]entities.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (u useCases) GetById(ctx context.Context, deviceId int64, userId int64) (entities.Device, error) {
	//TODO implement me
	panic("implement me")
}

func NewUseCases(deviceRepo device.Repository) UseCases {
	return &useCases{
		deviceRepo: deviceRepo,
	}
}
