package user

import (
	"base/domain/entities"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) GetByEmail(email string) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}
