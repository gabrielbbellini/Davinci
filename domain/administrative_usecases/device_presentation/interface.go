package device_presentation

import "context"

type UseCases interface {
	Relate(ctx context.Context, userId int64, deviceId int64, presentationId int64) error
	SetCurrentPresentation(ctx context.Context, userId int64, deviceId int64, presentationId int64) error
	GetCurrentPresentation(ctx context.Context, userId int64, deviceId int64) (int64, error)
}
