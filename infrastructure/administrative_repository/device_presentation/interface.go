package device_presentation

import "context"

type Repository interface {
	Relate(ctx context.Context, deviceId int64, presentationId int64) error
	SetCurrentPresentation(ctx context.Context, deviceId int64, presentationId int64) error
	GetCurrentPresentation(ctx context.Context, deviceId int64) (int64, error)
}
