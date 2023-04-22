package resolution

import (
	"base/domain/entities"
	"context"
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func (r repository) GetAll(ctx context.Context) ([]entities.Resolution, error) {
	var resolutions []entities.Resolution

	query := `
	SELECT id,
	       width,
	       height,
	       status_code, 
	FROM resolution as r
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error in [QueryContext]: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res entities.Resolution

		err = rows.Scan(
			&res.Id,
			&res.Width,
			&res.Height,
			&res.StatusCode,
		)

		if err != nil {
			log.Printf("Error in [Scan]: %v", err)
			return nil, err
		}

		resolutions = append(resolutions, res)
	}

	return resolutions, nil
}

func (r repository) GetById(ctx context.Context, id int64) (entities.Resolution, error) {
	var res entities.Resolution
	return res, nil
}

func NewResolutionRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
