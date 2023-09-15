package feature

import (
	"context"
	"eko-car/domain/logistik/constant"
	"eko-car/domain/logistik/model"
	Error "eko-car/domain/shared/error"
	"errors"
)

func (lf logistikFeature) AddProductFeature(ctx context.Context, request *model.AddProductRequest) (resp model.AddedProductResponse, err error) {

	// Added Bussiness logic here
	exist, err := lf.logistikRepo.CheckProductSKURepository(ctx, request.SKU)
	if err != nil {
		return
	} else if exist {
		err = Error.New(constant.ErrGeneral, constant.ErrSKUAlreadyExist, errors.New(request.SKU))
		return
	}

	product := model.Product{
		Name:  request.Name,
		SKU:   request.SKU,
		UOM:   request.UOM,
		Price: request.Price,
	}

	id, err := lf.logistikRepo.InsertProductRepository(ctx, product)
	if err != nil {
		return
	}

	resp = model.AddedProductResponse{
		Id:   id,
		Name: product.Name,
	}

	// userId := 1
	// Check Health sales
	// lf.queueService.PublishData(ctx, constant.CONSUMER_PRODUCT_INSERT_RABBITMQ, userId)

	return
}
