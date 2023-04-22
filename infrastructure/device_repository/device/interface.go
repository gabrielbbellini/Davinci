package device

import "context"

type Repository interface {
	get(ctx context.Context, userId int64, deviceId int64)
}
