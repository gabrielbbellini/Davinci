package presentation

import (
	"context"
	"davinci/domain/entities"
)

type Repository interface {
	// GetCurrentPresentation return the presentation which is playing.
	GetCurrentPresentation(ctx context.Context, id int64) (*entities.Presentation, error)
}
