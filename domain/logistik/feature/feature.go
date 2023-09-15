package feature

import (
	"context"
	"eko-car/domain/logistik/model"
	repository "eko-car/domain/logistik/repository"
	shared_model "eko-car/domain/shared/model"
)

type LogistikFeature interface {
	AddProductFeature(ctx context.Context, request *model.AddProductRequest) (response model.AddedProductResponse, err error)
	GetProductFeature(ctx context.Context, id string) (response model.Product, err error)
	GetProductListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (productList model.ProductLists, err error)
	DeleteProductFeature(ctx context.Context, id string) (response model.DeletedProductResponse, err error)
	UpdateProductFeature(ctx context.Context, id string, request *model.UpdateProductRequest) (response model.Product, err error)
	BulkCounterFeature(ctx context.Context) (err error)
	GetListsProductWithFilters(ctx context.Context, filter *shared_model.Filter) (productList model.ProductListsByFilter, err error)
}

type logistikFeature struct {
	logistikRepo repository.LogistikRepository
}

func NewLogistikFeature(logistikRepo repository.LogistikRepository) LogistikFeature {
	return &logistikFeature{
		logistikRepo: logistikRepo,
	}
}
