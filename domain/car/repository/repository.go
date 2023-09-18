package repository

import (
	"context"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/car/model"
	"eko-car/infrastructure/database"
)

type CarRepository interface {
	InsertCarRepository(ctx context.Context, product model.Car) (id int, err error)
	GetCarByIdRepository(ctx context.Context, id int) (product model.Car, err error)
	GetCarListsRepository(ctx context.Context, limit, offset int, sortby, search string, queryRequest shared_model.QueryRequest) (products []model.Car, err error)
	GetTotalCarRepository(ctx context.Context) (count int, err error)
	DeleteCarRepository(ctx context.Context, id int) (err error)
	UpdateCarRepository(ctx context.Context, id int, update *model.UpdateCarRequest) (err error)
	CheckCarIdRepository(ctx context.Context, id int) (exist bool, err error)
	GetTotalCarWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	GetTotalCarWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	GetCarListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (products []model.Car, err error)
}

type carRepository struct {
	Database *database.Database
}

func NewCarRepository(db *database.Database) CarRepository {
	return &carRepository{
		Database: db,
	}
}
