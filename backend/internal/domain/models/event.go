package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/shabohin/holiday.git/internal/domain/errs"
)

type Event struct {
	ID string `json:"id"`
}

func (c *Event) Validate() error {
	err := validation.ValidateStruct(
		c,
		validation.Field(&c.ID, is.UUID),
	)
	if err != nil {
		return errs.FromValidationError(err)
	}
	return nil
}

type EventFilter struct {
	IDs        []string `json:"ids"`
	PageSize   *uint64  `json:"page_size"`
	PageNumber *uint64  `json:"page_number"`
	OrderBy    []string `json:"order_by"`
}

func (c *EventFilter) Validate() error {
	err := validation.ValidateStruct(
		c,
		validation.Field(&c.IDs),
		validation.Field(&c.PageSize),
		validation.Field(&c.PageNumber),
		validation.Field(&c.OrderBy),
	)
	if err != nil {
		return errs.FromValidationError(err)
	}
	return nil
}

type EventCreate struct {
}

func (c *EventCreate) Validate() error {
	err := validation.ValidateStruct(
		c,
	)
	if err != nil {
		return errs.FromValidationError(err)
	}
	return nil
}

type EventUpdate struct {
	ID string `json:"id"`
}

func (c *EventUpdate) Validate() error {
	err := validation.ValidateStruct(
		c,
		validation.Field(&c.ID, validation.Required, is.UUID),
	)
	if err != nil {
		return errs.FromValidationError(err)
	}
	return nil
}
