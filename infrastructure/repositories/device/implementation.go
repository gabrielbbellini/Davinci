package device

import (
	"base/domain/entities"
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func (r repository) Create(ctx context.Context, device entities.Device) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, device entities.Device) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, device entities.Device) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetAll(ctx context.Context) ([]entities.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetById(ctx context.Context, id int) (entities.Device, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
