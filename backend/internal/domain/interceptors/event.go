package interceptors

import (
	"context"

	"github.com/shabohin/holiday.git/internal/domain/models"
)

//nolint: lll
//go:generate mockgen -destination mock/event_mock.go github.com/shabohin/holiday.git/internal/domain/interceptors EventInterceptor

type EventInterceptor interface {
	Get(ctx context.Context, id string, user *models.User) (*models.Event, error)
	List(ctx context.Context, filter *models.EventFilter, user *models.User) ([]*models.Event, error)
	Create(ctx context.Context, create *models.EventCreate, user *models.User) (*models.Event, error)
	Update(ctx context.Context, update *models.EventUpdate, user *models.User) (*models.Event, error)
	Delete(ctx context.Context, id string, user *models.User) error
}
