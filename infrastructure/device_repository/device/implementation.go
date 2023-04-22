package device

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func (r repository) get(ctx context.Context, userId int64, deviceId int64) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
