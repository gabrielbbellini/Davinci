package authorization

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	"davinci/settings"
	"davinci/view/http_error"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type repository struct {
	db       *sql.DB
	settings settings.Settings
}

func NewRepository(settings settings.Settings, db *sql.DB) Repository {
	return &repository{
		db:       db,
		settings: settings,
	}
}

func (r repository) Login(ctx context.Context, credential entities.Credential) (*entities.User, error) {
	user, err := r.getUserByEmail(ctx, credential.Email)
	if err != nil {
		log.Println("[Login] Error getUserByEmail", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Credential.Password), []byte(credential.Password))
	if err != nil {
		log.Println("[Login] Error bcrypt.CompareHashAndPassword", err)
		return nil, http_error.NewInternalServerError(http_error.ForbiddenMessage)
	}

	return user, err
}

func (r repository) getUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	//language=sql
	query := `
	SELECT id, 
	       name, 
	       email, 
	       password, 
	       id_role, 
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
		log.Println("[getByEmail] Error QueryRowContext", err)
		return nil, http_error.NewInternalServerError(err.Error())
	}

	return &user, err
}
