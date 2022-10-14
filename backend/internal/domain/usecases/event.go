package usecases

import (
	"context"

	"github.com/shabohin/holiday.git/internal/domain/models"
)

//nolint: lll
//go:generate mockgen -destination mock/event_mock.go github.com/shabohin/holiday.git/internal/domain/usecases EventUseCase

type EventUseCase interface {
	Get(ctx context.Context, id string) (*models.Event, error)
	List(ctx context.Context, filter *models.EventFilter) ([]*models.Event, error)
	Create(ctx context.Context, create *models.EventCreate) (*models.Event, error)
	Update(ctx context.Context, update *models.EventUpdate) (*models.Event, error)
	Delete(ctx context.Context, id string) error
}
