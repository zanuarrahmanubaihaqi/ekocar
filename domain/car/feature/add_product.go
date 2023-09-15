package feature

import (
	"context"
	"eko-car/domain/car/constant"
	"eko-car/domain/car/model"
	Error "eko-car/domain/shared/error"
	"errors"
)

func (lf carFeature) AddCarFeature(ctx context.Context, request *model.AddCarRequest) (resp model.AddedCarResponse, err error) {

	// Added Bussiness logic here
	exist, err := lf.carRepo.CheckCarSKURepository(ctx, request.SKU)
	if err != nil {
		return
	} else if exist {
		err = Error.New(constant.ErrGeneral, constant.ErrSKUAlreadyExist, errors.New(request.SKU))
		return
	}

	product := model.Car{
		Name:  request.Name,
		SKU:   request.SKU,
		UOM:   request.UOM,
		Price: request.Price,
	}

	id, err := lf.carRepo.InsertCarRepository(ctx, product)
	if err != nil {
		return
	}

	resp = model.AddedCarResponse{
		Id:   id,
		Name: product.Name,
	}

	// userId := 1
	// Check Health sales
	// lf.queueService.PublishData(ctx, constant.CONSUMER_PRODUCT_INSERT_RABBITMQ, userId)

	return
}
