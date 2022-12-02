package dto

import (
	"dont/hexagonal/errs"
	"dont/hexagonal/utils"
	"time"

	"github.com/google/uuid"
)

type ItemRequest struct {
	ID    uuid.UUID `json:"id" validate:"required,uuid"`
	Title string    `json:"title" validate:"required,lte=255"`
}

func (item ItemRequest) Validate() *errs.AppError {
	valid := utils.NewValidator()
	if err := valid.Struct(item); err != nil {
		// Return, if some fields are not valid.
		return errs.NewValidationError(err.Error())
	}
	return nil
}

type ItemResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
}
