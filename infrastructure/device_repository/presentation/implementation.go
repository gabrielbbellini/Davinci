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

func (r repository) GetAll(ctx context.Context, idUser int64, idResolution int64) ([]entities.Presentation, error) {
	query := `
	SELECT p.id,
	       p.name,
	       p.status_code, 
	       p.created_at, 
	       p.modified_at
	FROM presentation as p
	WHERE id_user = ? AND id_resolution = ?
	`

	result, err := r.db.QueryContext(ctx, query, idUser, idResolution)
	if err != nil {
		log.Println("[GetAll] error on QueryContext", err)
		return nil, err
	}
	defer result.Close()

	presentations := make([]entities.Presentation, 0)
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
			log.Println("[GetAll] error in Scan", err)
			return nil, err
		}

		presentations = append(presentations, presentation)
	}

	return presentations, nil
}

func (r repository) GetById(ctx context.Context, id int64, idUser int64) (*entities.Presentation, error) {
	query := `
	SELECT p.id,
	       p.name,
	       p.status_code, 
	       p.created_at, 
	       p.modified_at,
	       p.id_resolution,
	       r.width,
	       r.height
	FROM presentation as p
		INNER JOIN resolution r on p.id_resolution = r.id
	WHERE p.id = ? AND p.id_user = ?
	`

	queryPages := `
	SELECT p.id, 
	       p.id_presentation, 
	       p.timing, 
	       p.metadata, 
	       p.status_code, 
	       p.created_at, 
	       p.modified_at
	FROM page p
	WHERE id_presentation = ?
	`

	var presentation entities.Presentation
	var resolution entities.Resolution
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
		&resolution.Id,
		&resolution.Width,
		&resolution.Height,
	)
	if err != nil {
		log.Println("[GetById] error in QueryRowContext", err)
		return nil, err
	}

	result, err := r.db.QueryContext(ctx, queryPages, id)
	if err != nil {
		log.Println("[GetById] error in QueryContext", err)
		return nil, err
	}
	defer result.Close()

	var pages []entities.Page
	for result.Next() {
		var metadata string
		var page entities.Page

		err = result.Scan(
			&page.Id,
			&page.IdPresentation,
			&page.Timing,
			&metadata,
			&page.StatusCode,
			&page.CreatedAt,
			&page.ModifiedAt,
		)
		if err != nil {
			log.Println("[GetById] error in Scan", err)
			return nil, err
		}

		page.Metadata = metadata

		pages = append(pages, page)
	}

	presentation.Resolution = resolution
	presentation.Pages = pages

	return &presentation, nil
}

func NewPresentationRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
