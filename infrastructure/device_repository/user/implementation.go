package user

import (
	"base/domain/entities"
	"context"
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	//language=sql
	query := `
	SELECT id, 
	       name, 
	       email, 
	       password, 
	       roleId, 
	       status_code
	FROM user
	WHERE email = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Credential.Email,
		&user.Credential.Password,
		&user.Credential.RoleId,
		&user.StatusCode,
	)
	if err != nil {
		log.Println("[GetByEmail] Error Scan", err)
		return nil, err
	}

	return &user, nil
}
