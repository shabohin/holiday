package repositories

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"

	"github.com/shabohin/holiday.git/pkg/log"

	"github.com/shabohin/holiday.git/internal/domain/models"
	"github.com/shabohin/holiday.git/internal/domain/repositories"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/shabohin/holiday.git/internal/domain/errs"
)

type PostgresEventRepository struct {
	database *sqlx.DB
	logger   log.Logger
}

func NewPostgresEventRepository(database *sqlx.DB, logger log.Logger) repositories.EventRepository {
	return &PostgresEventRepository{database: database, logger: logger}
}

func (r *PostgresEventRepository) Create(ctx context.Context, event *models.Event) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	q := sq.Insert("public.events").
		Columns(). // TODO: add columns
		Values().  // TODO: add values
		Suffix("RETURNING \"id\"")
	query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
	if err := r.database.QueryRowxContext(ctx, query, args...).Scan(&event.ID); err != nil {
		e := errs.NewUnexpectedBehaviorError(err.Error())
		return e
	}
	return nil
}

func (r *PostgresEventRepository) Get(ctx context.Context, id string) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	event := &models.Event{}
	q := sq.Select("*").
		From("public.events").
		Where(sq.Eq{"id": id}).
		Limit(1)
	query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
	if err := r.database.GetContext(ctx, &event, query, args...); err != nil {
		e := errs.NewUnexpectedBehaviorError(err.Error())
		return nil, e
	}
	return event, nil
}

func (r *PostgresEventRepository) List(ctx context.Context, filter *models.EventFilter) ([]*models.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	var events []*models.Event
	const pageSize = 10
	q := sq.Select("*").
		From("public.events").
		Limit(pageSize) //
	// TODO: add filtering
	if filter.PageNumber != nil && *filter.PageNumber > 1 {
		q = q.Offset((*filter.PageNumber - 1) * *filter.PageSize)
	}
	if filter.PageSize != nil {
		q = q.Limit(*filter.PageSize)
	}
	if len(filter.OrderBy) > 0 {
		q = q.OrderBy(filter.OrderBy...)
	}
	query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
	if err := r.database.SelectContext(ctx, &events, query, args...); err != nil {
		e := errs.NewUnexpectedBehaviorError(err.Error())
		return nil, e
	}
	return events, nil
}

func (r *PostgresEventRepository) Update(ctx context.Context, event *models.Event) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	q := sq.Update("public.events").Where(sq.Eq{"id": event.ID}).Set("", "") // TODO: set values
	query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
	result, err := r.database.ExecContext(ctx, query, args...)
	if err != nil {
		e := errs.NewUnexpectedBehaviorError(err.Error())
		pgError, ok := err.(*pq.Error)
		if ok {
			switch pgError.Code {
			case "23505":
				e = errs.NewInvalidFormError()
				e.AddParam("phone", "The phone field has already been taken.")
			default:
				e = errs.NewUnexpectedBehaviorError(pgError.Detail)
			}
		}
		e.AddParam("event_id", fmt.Sprint(event.ID))
		return e
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return errs.NewUnexpectedBehaviorError(err.Error())
	}
	if affected == 0 {
		e := errs.NewEventNotFound()
		e.AddParam("event_id", fmt.Sprint(event.ID))
		return e
	}
	return nil
}

func (r *PostgresEventRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	q := sq.Delete("public.events").Where(sq.Eq{"id": id})
	query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
	result, err := r.database.ExecContext(ctx, query, args...)
	if err != nil {
		e := errs.NewUnexpectedBehaviorError(err.Error())
		e.AddParam("event_id", fmt.Sprint(id))
		return e
	}
	affected, err := result.RowsAffected()
	if err != nil {
		e := errs.NewUnexpectedBehaviorError(err.Error())
		e.AddParam("event_id", fmt.Sprint(id))
		return e
	}
	if affected == 0 {
		e := errs.NewEventNotFound()
		e.AddParam("event_id", fmt.Sprint(id))
		return e
	}
	return nil
}
