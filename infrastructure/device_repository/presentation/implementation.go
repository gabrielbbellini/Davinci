package presentation

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	"encoding/json"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewPresentationRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) GetCurrentPresentation(ctx context.Context, deviceId int64) (*entities.Presentation, error) {
	presentation, err := r.getCurrentPresentation(ctx, deviceId)
	if err != nil {
		log.Println("[GetCurrentPresentation] Error getCurrentPresentation", err)
		return nil, err
	}

	presentation.Pages, err = r.getPresentationPages(ctx, presentation.Id)
	if err != nil {
		log.Println("[GetCurrentPresentation] Error getPresentationPages", err)
		return nil, err
	}

	return presentation, nil
}

func (r repository) getCurrentPresentation(ctx context.Context, deviceId int64) (*entities.Presentation, error) {
	//language=sql
	query := `
	SELECT p.id, 
	       p.id_resolution, 
	       p.id_orientation, 
	       p.name, 
	       p.status_code, 
	       p.created_at, 
	       p.modified_at
	FROM presentation p
	INNER JOIN device_presentation dp ON 
	    p.id = dp.id_presentation AND
	    dp.id_device = ? AND
	    dp.is_playing AND
	    dp.status_code = ?
	WHERE p.status_code = ?`

	var presentation entities.Presentation
	err := r.db.QueryRowContext(ctx, query, deviceId, entities.StatusOk, entities.StatusOk).Scan(
		&presentation.Id,
		&presentation.ResolutionId,
		&presentation.Orientation,
		&presentation.Name,
		&presentation.StatusCode,
		&presentation.CreatedAt,
		&presentation.ModifiedAt,
	)
	if err != nil {
		log.Println("[getCurrentPresentation] Error Scan", err)
		return nil, err
	}

	return &presentation, nil
}

func (r repository) getPresentationPages(ctx context.Context, presentationId int64) ([]entities.Page, error) {
	//language=sql
	query := `
	SELECT id, 
	       id_presentation, 
	       duration, 
	       component, 
	       status_code, 
	       created_at, 
	       modified_at
	FROM page
	WHERE id_presentation = ? AND status_code = ?
	`

	rows, err := r.db.QueryContext(ctx, query, presentationId, entities.StatusOk)
	if err != nil {
		log.Println("[getPresentationPages] Error QueryContext", err)
		return nil, err
	}
	defer rows.Close()

	var pages []entities.Page
	for rows.Next() {
		var page entities.Page
		var componentString string
		err = rows.Scan(
			&page.Id,
			&page.PresentationId,
			&page.Duration,
			&componentString,
			&page.StatusCode,
			&page.CreatedAt,
			&page.ModifiedAt,
		)
		if err != nil {
			log.Println("[getPresentationPages] Error Scan", err)
			return nil, err
		}

		err = json.Unmarshal([]byte(componentString), &page.Component)
		if err != nil {
			log.Println("[getPresentationPages] Error Unmarshal", err)
			return nil, err
		}

		pages = append(pages, page)
	}

	return pages, err
}
