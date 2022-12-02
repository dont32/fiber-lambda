package domain

import (
	"dont/hexagonal/dto"
	"dont/hexagonal/errs"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Title     string    `db:"title"`
}

type ItemRepository interface {
	GetByID(uuid.UUID) (*Item, *errs.AppError)
	AddNew(Item) (*Item, *errs.AppError)
}

func (d Item) ToDto() dto.ItemResponse {
	return dto.ItemResponse{
		ID:        d.ID,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		Title:     d.Title,
	}
}

// type Tabler interface {
// 	TableName() string
// }

// // TableName overrides the table name used by User to `profiles`
// func (Item) TableName() string {
// 	return "Item"
// }
