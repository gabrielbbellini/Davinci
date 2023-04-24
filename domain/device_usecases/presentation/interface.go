package presentation

import (
	"context"
	"davinci/domain/entities"
)

type UseCases interface {
	// GetCurrentPresentation get the presentation which the status on device is playing.
	GetCurrentPresentation(ctx context.Context, id int64) (*entities.Presentation, error)
}
