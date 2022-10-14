package usecases

import (
	"context"

	"github.com/shabohin/holiday.git/internal/domain/models"
	"github.com/shabohin/holiday.git/internal/domain/repositories"
	"github.com/shabohin/holiday.git/internal/domain/usecases"

	"github.com/shabohin/holiday.git/pkg/log"
)

//nolint: lll
//go:generate mockgen -destination mock/event_mock.go github.com/shabohin/holiday.git/internal/usecases EventUseCase

type EventUseCase struct {
	eventRepository repositories.EventRepository
	logger          log.Logger
}

func NewEventUseCase(
	eventRepository repositories.EventRepository,
	logger log.Logger,
) usecases.EventUseCase {
	return &EventUseCase{
		eventRepository: eventRepository,
		logger:          logger,
	}
}

func (u *EventUseCase) Get(ctx context.Context, id string) (*models.Event, error) {
	qr, err := u.eventRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return qr, nil
}

func (u *EventUseCase) List(ctx context.Context, filter *models.EventFilter) ([]*models.Event, error) {
	qrs, err := u.eventRepository.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return qrs, nil
}

func (u *EventUseCase) Create(ctx context.Context, create *models.EventCreate) (*models.Event, error) {
	if err := create.Validate(); err != nil {
		return nil, err
	}
	event := &models.Event{
		ID: "",
	}

	if err := u.eventRepository.Create(ctx, event); err != nil {
		return nil, err
	}
	return event, nil
}

func (u *EventUseCase) Update(ctx context.Context, update *models.EventUpdate) (*models.Event, error) {
	if err := update.Validate(); err != nil {
		return nil, err
	}
	event, err := u.eventRepository.Get(ctx, update.ID)
	if err != nil {
		return nil, err
	}
	if err := u.eventRepository.Update(ctx, event); err != nil {
		return nil, err
	}
	return event, nil
}

func (u *EventUseCase) Delete(ctx context.Context, id string) error {
	if err := u.eventRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
