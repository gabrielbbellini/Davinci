package device

import (
	"base/domain/entities"
	"context"
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(ctx context.Context, device entities.Device, userId int64) error {
	query := `
	INSERT INTO device (name, id_resolution, id_orientation, id_user) 
	VALUES (?,?,?,?)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		device.Name,
		device.Resolution.Id,
		device.Orientation,
		userId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Update(ctx context.Context, device entities.Device, userId int64) error {
	command := `
	UPDATE device
	SET name = ?, 
	    id_resolution = ?, 
	    id_orientation = ?
	WHERE id = ? AND 
	      id_user = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		command,
		device.Name,
		device.Resolution.Id,
		device.Orientation,
		device.Id,
		userId,
	)
	if err != nil {
		log.Println("[Update] Error ExecContext", err)
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, device entities.Device, userId int64) error {
	query := `
	DELETE FROM device d
	WHERE id = ? AND id_user = ?
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, device.Id, userId)
	if err != nil {
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
	       d.created_at, 
	       d.modified_at
	FROM device as d
	WHERE id_user = ?
	`
	rows, err := r.db.QueryContext(ctx, query, userId)
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
			&device.CreatedAt,
			&device.ModifiedAt,
		)
		if err != nil {
			log.Println("[GetAll] Error Scan", err)
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (r repository) GetById(ctx context.Context, id int64, userId int64) (entities.Device, error) {
	query := `
	SELECT d.id,
	       d.name,
	       d.id_orientation,
	       d.status_code, 
	       d.created_at, 
	       d.modified_at,
	       d.id_resolution,
	       r.width,
	       r.height
	FROM device as d
		INNER JOIN resolution r on d.id_resolution = r.id
	WHERE d.id = ? AND d.id_user = ?
	`
	var device entities.Device

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
		userId,
	).Scan(
		&device.Id,
		&device.Name,
		&device.Orientation,
		&device.StatusCode,
		&device.CreatedAt,
		&device.ModifiedAt,
		&device.Resolution.Id,
		&device.Resolution.Width,
		&device.Resolution.Height,
	)

	if err != nil {
		return device, err
	}

	return device, nil
}
