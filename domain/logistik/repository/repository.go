package repository

import (
	"eko-car/domain/logistik/model"
	shared_model "eko-car/domain/shared/model"
	"eko-car/infrastructure/database"
	"context"
	"database/sql"
)

type LogistikRepository interface {
	InsertProductRepository(ctx context.Context, product model.Product) (id int, err error)
	GetProductBySKURepository(ctx context.Context, sku string) (product model.Product, err error)
	GetProductByIdRepository(ctx context.Context, id int) (product model.Product, err error)
	GetProductListsRepository(ctx context.Context, limit, offset int, sortby, search string) (products []model.Product, err error)
	GetTotalProductRepository(ctx context.Context) (count int, err error)
	DeleteProductRepository(ctx context.Context, id int) (err error)
	UpdateProductRepository(ctx context.Context, id int, update *model.UpdateProductRequest) (err error)
	CheckProductIdRepository(ctx context.Context, id int) (exist bool, err error)
	CheckProductSKURepository(ctx context.Context, sku string) (exist bool, err error)
	GetTotalProductWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	BulkInsertCounterRepository(ctx context.Context, size int) (err error)
	GetLastCounterRepository(ctx context.Context) (number string, err error)
	GetDocNumberRangeRepository(ctx context.Context) (data model.NumberRange, err error)
	GetAndUpdateNumberNextRepository(ctx context.Context, tx *sql.Tx) (number int, err error)
	GetTotalProductWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	GetProductListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (products []model.Product, err error)
}

type logistikRepository struct {
	Database *database.Database
}

func NewLogistikRepository(db *database.Database) LogistikRepository {
	return &logistikRepository{
		Database: db,
	}
}
