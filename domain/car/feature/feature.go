package feature

import (
	"context"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/car/model"
	repository "eko-car/domain/car/repository"
)

type CarFeature interface {
	AddCarFeature(ctx context.Context, request *model.AddCarRequest) (response model.AddedCarResponse, err error)
	GetCarFeature(ctx context.Context, id int) (response model.Car, err error)
	GetCarListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (productList model.CarLists, err error)
	DeleteCarFeature(ctx context.Context, id int) (response model.DeletedCarResponse, err error)
	UpdateCarFeature(ctx context.Context, id int, request *model.UpdateCarRequest) (response model.Car, err error)
	GetListsCarWithFilters(ctx context.Context, filter *shared_model.Filter) (productList model.CarListsByFilter, err error)
}

type carFeature struct {
	carRepo repository.CarRepository
}

func NewCarFeature(carRepo repository.CarRepository) CarFeature {
	return &carFeature{
		carRepo: carRepo,
	}
}
