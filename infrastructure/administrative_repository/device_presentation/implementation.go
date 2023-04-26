package device_presentation

import (
	"context"
	"database/sql"
	"davinci/settings"
	"log"
)

type repository struct {
	db       *sql.DB
	settings settings.Settings
}

func NewRepository(settings settings.Settings, db *sql.DB) Repository {
	return &repository{db: db, settings: settings}
}

func (r repository) Relate(ctx context.Context, deviceId int64, presentationId int64) error {
	//language=sql
	query := `
	INSERT INTO device_presentation (id_device, id_presentation) 
	VALUES (?, ?)`

	_, err := r.db.ExecContext(ctx, query, deviceId, presentationId)
	if err != nil {
		log.Println("[Relate] Error ExecContext", err)
		return err
	}

	return nil
}

func (r repository) SetCurrentPresentation(ctx context.Context, deviceId int64, presentationId int64) error {
	//language=sql
	query := `
	UPDATE device_presentation 
	SET is_playing = 1
	WHERE id_device = ? AND 
	      id_presentation = ?`

	_, err := r.db.ExecContext(ctx, query, deviceId, presentationId)
	if err != nil {
		log.Println("[SetCurrentPresentation] Error ExecContext", err)
		return err
	}

	return nil
}

func (r repository) GetCurrentPresentation(ctx context.Context, deviceId int64) (int64, error) {
	//language=sql
	query := `
	SELECT id_presentation
	FROM device_presentation
	WHERE id_device = ? AND 
	      is_playing`

	var presentationId int64
	err := r.db.QueryRowContext(ctx, query, deviceId).Scan(&presentationId)
	if err != nil {
		log.Println("[GetCurrentPresentation] Scan", err)
		return 0, err
	}

	return presentationId, nil
}
