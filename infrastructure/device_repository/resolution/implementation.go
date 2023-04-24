package resolution

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewResolutionRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) GetAll(ctx context.Context) ([]entities.Resolution, error) {
	query := `
	SELECT id,
	       width,
	       height,
	       status_code 
	FROM resolution
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("[GetAll] Error QueryContext", err)
		return nil, err
	}
	defer rows.Close()

	var resolutions []entities.Resolution
	for rows.Next() {
		var res entities.Resolution
		err = rows.Scan(
			&res.Id,
			&res.Width,
			&res.Height,
			&res.StatusCode,
		)
		if err != nil {
			log.Println("[GetAll] Error Scan", err)
			return nil, err
		}
		resolutions = append(resolutions, res)
	}

	return resolutions, nil
}

func (r repository) GetById(ctx context.Context, resolutionId int64) (*entities.Resolution, error) {
	query := `
	SELECT id,
	       width,
	       height,
	       status_code
	FROM resolution
	WHERE id = ?`

	var resolution entities.Resolution
	err := r.db.QueryRowContext(ctx, query, resolutionId).Scan(
		&resolution.Id,
		&resolution.Width,
		&resolution.Height,
		&resolution.StatusCode,
	)
	if err != nil {
		log.Println("[GetAll] Error Scan", err)
		return nil, err
	}

	return &resolution, nil
}
