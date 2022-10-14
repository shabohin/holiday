package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/shabohin/holiday.git/internal/domain/errs"
)

type User struct {
	ID string `json:"id"`
}

func (c *User) Validate() error {
	err := validation.ValidateStruct(
		c,
		validation.Field(&c.ID, is.UUID),
	)
	if err != nil {
		return errs.FromValidationError(err)
	}
	return nil
}

type UserFilter struct {
	IDs        []string `json:"ids"`
	PageSize   *uint64  `json:"page_size"`
	PageNumber *uint64  `json:"page_number"`
	OrderBy    []string `json:"order_by"`
}

func (c *UserFilter) Validate() error {
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

type UserCreate struct {
}

func (c *UserCreate) Validate() error {
	err := validation.ValidateStruct(
		c,
	)
	if err != nil {
		return errs.FromValidationError(err)
	}
	return nil
}

type UserUpdate struct {
	ID string `json:"id"`
}

func (c *UserUpdate) Validate() error {
	err := validation.ValidateStruct(
		c,
		validation.Field(&c.ID, validation.Required, is.UUID),
	)
	if err != nil {
		return errs.FromValidationError(err)
	}
	return nil
}
