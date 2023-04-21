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
	SELECT r.id,
	       r.width,
	       r.height,
	       r.status_code, 
	       r.created_at, 
	       r.modified_at
	FROM resolution as r
	`
	result, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error in [QueryContext]: %v", err)
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var res entities.Resolution

		err = result.Scan(
			&res.Id,
			&res.Width,
			&res.Height,
			&res.StatusCode,
			&res.CreatedAt,
			&res.ModifiedAt,
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
