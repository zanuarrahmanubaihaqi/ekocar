package repository

import (
	"eko-car/domain/car/model"
	shared_model "eko-car/domain/shared/model"
	"eko-car/infrastructure/database"
	"context"
	"database/sql"
)

type CarRepository interface {
	InsertCarRepository(ctx context.Context, product model.Car) (id int, err error)
	GetCarBySKURepository(ctx context.Context, sku string) (product model.Car, err error)
	GetCarByIdRepository(ctx context.Context, id int) (product model.Car, err error)
	GetCarListsRepository(ctx context.Context, limit, offset int, sortby, search string) (products []model.Car, err error)
	GetTotalCarRepository(ctx context.Context) (count int, err error)
	DeleteCarRepository(ctx context.Context, id int) (err error)
	UpdateCarRepository(ctx context.Context, id int, update *model.UpdateCarRequest) (err error)
	CheckCarIdRepository(ctx context.Context, id int) (exist bool, err error)
	CheckCarSKURepository(ctx context.Context, sku string) (exist bool, err error)
	GetTotalCarWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	BulkInsertCounterRepository(ctx context.Context, size int) (err error)
	GetLastCounterRepository(ctx context.Context) (number string, err error)
	GetDocNumberRangeRepository(ctx context.Context) (data model.NumberRange, err error)
	GetAndUpdateNumberNextRepository(ctx context.Context, tx *sql.Tx) (number int, err error)
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
