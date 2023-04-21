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

func (r repository) Create(ctx context.Context, device entities.Device) error {
	query := `
	INSERT INTO device (name, id_resolution, id_orientation) 
	VALUES (?,?,?)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, device.Name, device.Resolution.Id, device.Orientation)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Update(ctx context.Context, device entities.Device) error {
	query := `
	UPDATE device d
	SET
	    d.name = ?, 
	    d.id_resolution = ?, 
	    d.id_orientation = ?
	WHERE id = ?
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, device.Name, device.Resolution.Id, device.Orientation, device.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, device entities.Device) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetAll(ctx context.Context) ([]entities.Device, error) {
	devices := make([]entities.Device, 0)

	query := `
	SELECT d.id,
	       d.name,
	       d.id_orientation,
	       d.status_code, 
	       d.created_at, 
	       d.modified_at
	FROM device as d
	`
	result, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error in [QueryContext]: %v", err)
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var dev entities.Device

		err = result.Scan(
			&dev.Id,
			&dev.Name,
			&dev.Orientation,
			&dev.StatusCode,
			&dev.CreatedAt,
			&dev.ModifiedAt,
		)
		if err != nil {
			log.Printf("Error in [Scan]: %v", err)
			return nil, err
		}

		devices = append(devices, dev)
	}

	return devices, nil
}

func (r repository) GetById(ctx context.Context, id int64) (entities.Device, error) {
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
	WHERE d.id = ?
	`
	var dev entities.Device
	var res entities.Resolution

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&dev.Id,
		&dev.Name,
		&dev.Orientation,
		&dev.StatusCode,
		&dev.CreatedAt,
		&dev.ModifiedAt,
		&res.Id,
		&res.Width,
		&res.Height,
	)

	if err != nil {
		return dev, err
	}

	dev.Resolution = &res

	return dev, nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
