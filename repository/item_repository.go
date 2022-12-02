package repository

import (
	"dont/hexagonal/domain"
	"dont/hexagonal/errs"
	"net"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemsRepositoryDb struct {
	dbClient *gorm.DB
}

func (db ItemsRepositoryDb) GetByID(id uuid.UUID) (*domain.Item, *errs.AppError) {
	item := domain.Item{}
	item.ID = id
	rs := db.dbClient.First(&item)
	if rs.Error != nil {
		//errors.Is(rs.Error, gorm.Err)
		err, _ := rs.Error.(*net.OpError)
		if err != nil {
			return nil, errs.NewUnexpectedError("database unexpected")
		}

		return nil, errs.NewUnexpectedError(rs.Error.Error())
	}
	return &item, nil
}

func (db ItemsRepositoryDb) AddNew(item domain.Item) (*domain.Item, *errs.AppError) {
	rs := db.dbClient.Create(&item)
	if rs.Error != nil {
		return nil, errs.NewUnexpectedError("database unexpected")
	}
	return &item, nil
}

func NewItemRepositoryDb(client *gorm.DB) ItemsRepositoryDb {
	return ItemsRepositoryDb{dbClient: client}
}
