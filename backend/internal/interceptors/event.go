package interceptors

import (
	"context"

	"github.com/shabohin/holiday.git/internal/domain/interceptors"
	"github.com/shabohin/holiday.git/internal/domain/models"
	"github.com/shabohin/holiday.git/internal/domain/usecases"

	"github.com/shabohin/holiday.git/pkg/log"
)

//nolint: lll
//go:generate mockgen -destination mock/event_mock.go github.com/shabohin/holiday.git/internal/interceptors EventInterceptor

type EventInterceptor struct {
	eventUseCase usecases.EventUseCase
	logger       log.Logger
}

func NewEventInterceptor(
	eventUseCase usecases.EventUseCase,
	logger log.Logger,
) interceptors.EventInterceptor {
	return &EventInterceptor{
		eventUseCase: eventUseCase,
		logger:       logger,
	}
}

func (i *EventInterceptor) Get(ctx context.Context, id string, user *models.User) (*models.Event, error) {
	event, err := i.eventUseCase.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (i *EventInterceptor) List(
	ctx context.Context,
	filter *models.EventFilter,
	user *models.User,
) ([]*models.Event, error) {
	events, err := i.eventUseCase.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (i *EventInterceptor) Create(
	ctx context.Context,
	create *models.EventCreate,
	user *models.User,
) (*models.Event, error) {
	event, err := i.eventUseCase.Create(ctx, create)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (i *EventInterceptor) Update(
	ctx context.Context,
	update *models.EventUpdate,
	user *models.User,
) (*models.Event, error) {
	event, err := i.eventUseCase.Update(ctx, update)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (i *EventInterceptor) Delete(ctx context.Context, id string, user *models.User) error {
	if err := i.eventUseCase.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
