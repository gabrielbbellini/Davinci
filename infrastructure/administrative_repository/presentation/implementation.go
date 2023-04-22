package presentation

import (
	"base/domain/entities"
	"context"
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func (r repository) Create(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	query := `
	INSERT INTO presentation (name,id_user) 
	VALUES (?,?)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		presentation.Name,
		idUser,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Update(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	query := `
	UPDATE presentation d
	SET
	    d.name = ?
	WHERE id = ? AND id_user = ?
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, presentation.Name, presentation.Id, presentation.Id, idUser)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	query := `
	DELETE FROM presentation d
	WHERE id = ? AND id_user = ?
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, presentation.Id, idUser)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetAll(ctx context.Context, idUser int64) ([]entities.Presentation, error) {
	presentations := make([]entities.Presentation, 0)

	query := `
	SELECT d.id,
	       d.name,
	       d.status_code, 
	       d.created_at, 
	       d.modified_at
	FROM presentation as d
	WHERE id_user = ?
	`
	result, err := r.db.QueryContext(ctx, query, idUser)
	if err != nil {
		log.Printf("Error in [QueryContext]: %v", err)
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var presentation entities.Presentation

		err = result.Scan(
			&presentation.Id,
			&presentation.Name,
			&presentation.StatusCode,
			&presentation.CreatedAt,
			&presentation.ModifiedAt,
		)
		if err != nil {
			log.Printf("Error in [Scan]: %v", err)
			return nil, err
		}

		presentations = append(presentations, presentation)
	}

	return presentations, nil
}

func (r repository) GetById(ctx context.Context, id int64, idUser int64) (entities.Presentation, error) {
	query := `
	SELECT p.id,
	       p.id_user,
	       p.name,
	       p.status_code, 
	       p.created_at, 
	       p.modified_at
	FROM presentation as p
	WHERE p.id = ? AND p.id_user = ?
	`
	var presentation entities.Presentation

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
		idUser,
	).Scan(
		&presentation.Id,
		&presentation.Name,
		&presentation.StatusCode,
		&presentation.CreatedAt,
		&presentation.ModifiedAt,
	)

	if err != nil {
		return presentation, err
	}

	return presentation, nil
}

func NewPresentationRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
