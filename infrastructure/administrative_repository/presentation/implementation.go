package presentation

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	"davinci/settings"
	"encoding/json"
	"log"
)

type repository struct {
	db       *sql.DB
	settings settings.Settings
}

func NewPresentationRepository(settings settings.Settings, db *sql.DB) Repository {
	return &repository{
		db:       db,
		settings: settings,
	}
}

func (r repository) Create(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	queryPresentation := `
	INSERT INTO presentation (name, id_user, id_resolution)
	VALUES (?,?,?)
	`

	queryPages := `
	INSERT INTO page (id_presentation, duration, component) 
	VALUES (?,?,?)
	`

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("[Create] error in Begin", err)
		return err
	}

	stmt, err := tx.PrepareContext(ctx, queryPresentation)
	if err != nil {
		log.Println("[Create] error in PrepareContext Presentation", err)
		return err
	}
	defer stmt.Close()

	stmtPages, err := tx.PrepareContext(ctx, queryPages)
	if err != nil {
		log.Println("[Create] error in PrepareContext Presentation", err)
		return err
	}
	defer stmtPages.Close()

	result, err := stmt.ExecContext(
		ctx,
		presentation.Name,
		idUser,
	)
	if err != nil {
		log.Println("[Create] error in ExecContext Presentation", err)
		_ = tx.Rollback()
		return err
	}

	idPresentation, err := result.LastInsertId()

	for _, item := range presentation.Pages {

		var b []byte

		switch item.Metadata.(type) {
		case []interface{}:
			b, err = json.Marshal(item.Metadata)
			if err != nil {
				log.Printf("Unsupported struct on param metadata")
				_ = tx.Rollback()
				return err
			}
		default:
			log.Printf("Unsupported struct on param metadata")
			b = []byte("")
		}
		_, err = stmtPages.ExecContext(
			ctx,
			idPresentation,
			item.Timing,
			string(b),
		)
		if err != nil {
			log.Println("[Create] error in ExecContext Presentation", err)
			_ = tx.Rollback()
			return err
		}
	}

	_ = tx.Commit()
	return nil
}

func (r repository) Update(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	query := `
	UPDATE presentation p
	SET
	    p.name = ?,
	    p.id_resolution = ?,
	    p.status_code = ?
	WHERE p.id = ? AND p.id_user = ?
	`

	deletePages := `
	DELETE FROM page p
	WHERE id_presentation = ?;
	`

	queryPages := `
	INSERT INTO page (id_presentation, duration, component) 
	VALUES (?,?,?)
	`

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("[Update] error in Begin tx", err)
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		_ = tx.Rollback()
		log.Println("[Update] error in PrepareContext Presentation", err)
		return err
	}
	defer stmt.Close()

	stmtPage, err := tx.PrepareContext(ctx, queryPages)
	if err != nil {
		_ = tx.Rollback()
		log.Println("[Update] error in PrepareContext Pages", err)
		return err
	}
	defer stmt.Close()

	_, err = tx.ExecContext(ctx, deletePages, presentation.Id)
	if err != nil {
		_ = tx.Rollback()
		log.Println("[Update] error in Delete pages", err)
		return err
	}

	_, err = stmt.ExecContext(
		ctx,
		presentation.Name,
		presentation.StatusCode,
		presentation.Resolution.Id,
		presentation.Id,
		idUser,
	)
	if err != nil {
		_ = tx.Rollback()
		log.Println("[Update] error in ExecContext Presentation", err)
		return err
	}

	for _, item := range presentation.Pages {
		var b []byte

		switch item.Metadata.(type) {
		case []interface{}:
			b, err = json.Marshal(item.Metadata)
			if err != nil {
				log.Printf("Unsupported struct on param metadata")
				_ = tx.Rollback()
				return err
			}
		default:
			log.Printf("Unsupported struct on param metadata")
			b = []byte("")
		}

		_, err = stmtPage.ExecContext(ctx, presentation.Id, item.Timing, string(b))
		if err != nil {
			log.Printf("[Update] error in ExecContext Pages")
			_ = tx.Rollback()
			return err
		}
	}

	_ = tx.Commit()
	return nil
}

func (r repository) Delete(ctx context.Context, presentation entities.Presentation, idUser int64) error {
	queryPresentation := `
	DELETE FROM presentation d
	WHERE id = ? AND id_user = ?
	`

	queryPages := `
	DELETE FROM page p
	WHERE id_presentation = ?;
	`

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("[Delete] error in Begin", err)
		return err
	}

	_, err = tx.ExecContext(ctx, queryPresentation, presentation.Id, idUser)
	if err != nil {
		log.Println("[Delete] error in ExecContext Presentation", err)
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, queryPages, presentation.Id)
	if err != nil {
		log.Println("[Delete] error in ExecContext Pages", err)
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()

	return nil
}

func (r repository) GetAll(ctx context.Context, idUser int64) ([]entities.Presentation, error) {
	query := `
	SELECT p.id,
	       p.name,
	       p.status_code, 
	       p.created_at, 
	       p.modified_at
	FROM presentation as p
	WHERE id_user = ?
	`

	rows, err := r.db.QueryContext(ctx, query, idUser)
	if err != nil {
		log.Println("[GetAll] error on QueryContext", err)
		return nil, err
	}
	defer rows.Close()

	var presentations []entities.Presentation
	for rows.Next() {
		var presentation entities.Presentation
		err = rows.Scan(
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
	       p.duration, 
	       p.component, 
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
