package service

import (
	"dont/hexagonal/domain"
	"dont/hexagonal/dto"
	"dont/hexagonal/errs"
)

type ItemService interface {
	GetByID(dto.ItemRequest) (*dto.ItemResponse, *errs.AppError)
	AddNew(dto.ItemRequest) (*dto.ItemResponse, *errs.AppError)
}

type DefaultItemService struct {
	repo domain.ItemRepository
}

func (s DefaultItemService) GetByID(req dto.ItemRequest) (*dto.ItemResponse, *errs.AppError) {
	// err := req.Validate()
	// if err != nil {
	// 	return nil, err
	// }
	item := domain.Item{
		ID: req.ID,
	}
	res_domain, err := s.repo.GetByID(item.ID)
	if err != nil {
		return nil, err
	}
	res := res_domain.ToDto()
	return &res, nil
}

func (s DefaultItemService) AddNew(req dto.ItemRequest) (*dto.ItemResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	item := domain.Item{
		ID:    req.ID,
		Title: req.Title,
	}
	res_domain, err := s.repo.AddNew(item)
	if err != nil {
		return nil, err
	}
	res := res_domain.ToDto()
	return &res, nil
}

func NewItemService(repo domain.ItemRepository) DefaultItemService {
	return DefaultItemService{repo: repo}
}
