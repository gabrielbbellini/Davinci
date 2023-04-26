package device

import (
	"context"
	"database/sql"
	"davinci/domain/entities"
	"davinci/settings"
	"errors"
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

func (r repository) Create(ctx context.Context, device entities.Device, userId int64) error {
	query := `
	INSERT INTO device (name, id_resolution, id_orientation, id_user) 
	VALUES (?,?,?,?)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println("[Update] Error PrepareContext", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		device.Name,
		device.ResolutionId,
		device.Orientation,
		userId,
	)
	if err != nil {
		log.Println("[Create] Error ExecContext", err)
		return err
	}

	return nil
}

func (r repository) Update(ctx context.Context, deviceId int64, device entities.Device, userId int64) error {
	query := `
	UPDATE device
	SET name = ?, 
	    id_resolution = ?, 
	    id_orientation = ?
	WHERE id = ? AND 
	      id_user = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		device.Name,
		device.ResolutionId,
		device.Orientation,
		deviceId,
		userId,
	)
	if err != nil {
		log.Println("[Update] Error ExecContext", err)
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, deviceId int64, userId int64) error {
	//language=sql
	query := `
	UPDATE device 
	SET status_code = ?
	WHERE id = ? 
	  AND id_user = ?
	`

	_, err := r.db.ExecContext(ctx, query, entities.StatusDeleted, deviceId, userId)
	if err != nil {
		log.Println("[Delete] Error ExecContext", err)
		return err
	}

	return nil
}

func (r repository) GetAll(ctx context.Context, userId int64) ([]entities.Device, error) {
	query := `
	SELECT d.id,
	       d.name,
	       d.id_orientation,
	       d.status_code,
		   d.id_resolution,
		   d.created_at,
		   d.modified_at
	FROM device d
	WHERE id_user = ? AND 
	      d.status_code = ?
	`
	rows, err := r.db.QueryContext(ctx, query, userId, entities.StatusOk)
	if err != nil {
		log.Println("[GetAll] Error QueryContext", err)
		return nil, err
	}
	defer rows.Close()

	var devices []entities.Device
	for rows.Next() {
		var device entities.Device

		err = rows.Scan(
			&device.Id,
			&device.Name,
			&device.Orientation,
			&device.StatusCode,
			&device.ResolutionId,
			&device.CreatedAt,
			&device.ModifiedAt,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			log.Println("[GetAll] Error Scan", err)
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (r repository) GetById(ctx context.Context, deviceId int64, userId int64) (*entities.Device, error) {
	//language=sql
	query := `
	SELECT id,
	       name,
	       id_orientation,
	       status_code, 
	       id_resolution,
	       created_at,
	       modified_at
	FROM device
	WHERE id = ? AND 
	      id_user = ? AND
	      status_code = ?
	`

	var device entities.Device
	err := r.db.QueryRowContext(
		ctx,
		query,
		deviceId,
		userId,
		entities.StatusOk,
	).Scan(
		&device.Id,
		&device.Name,
		&device.Orientation,
		&device.StatusCode,
		&device.ResolutionId,
		&device.CreatedAt,
		&device.ModifiedAt,
	)

	if err != nil {
		log.Println("[GetById] Error Scan", err)
		return nil, err
	}

	return &device, nil
}

func (r repository) GetDeviceByName(ctx context.Context, deviceName string, userId int64) (*entities.Device, error) {
	//language=sql
	query := `
	SELECT d.id,
	       d.name,
	       d.id_orientation,
	       d.status_code, 
	       d.id_resolution
	FROM device d
	WHERE d.name = ? AND 
	      d.id_user = ? AND
	      d.status_code = ?`

	var device entities.Device
	err := r.db.QueryRowContext(
		ctx,
		query,
		deviceName,
		userId,
		entities.StatusOk,
	).Scan(
		&device.Id,
		&device.Name,
		&device.Orientation,
		&device.StatusCode,
		&device.ResolutionId,
	)

	if err != nil {
		log.Println("[GetDeviceByName] Error Scan", err)
		return nil, err
	}

	return &device, nil
}
