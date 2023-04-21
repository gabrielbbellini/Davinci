package user

import (
	"base/domain/entities"
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func (r repository) Create(ctx context.Context, user entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Get(ctx context.Context, id int64) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetAll(ctx context.Context) ([]entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, user entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
