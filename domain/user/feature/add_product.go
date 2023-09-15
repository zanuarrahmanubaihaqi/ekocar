package feature

import (
	"context"
	"eko-car/domain/user/constant"
	"eko-car/domain/user/model"
	Error "eko-car/domain/shared/error"
	"errors"
)

func (uf userFeature) AddUserFeature(ctx context.Context, request *model.AddUserRequest) (resp model.AddedUserResponse, err error) {

	// Added Bussiness logic here
	exist, err := uf.userRepo.CheckUserSKURepository(ctx, request.SKU)
	if err != nil {
		return
	} else if exist {
		err = Error.New(constant.ErrGeneral, constant.ErrSKUAlreadyExist, errors.New(request.SKU))
		return
	}

	product := model.User{
		Name:  request.Name,
		SKU:   request.SKU,
		UOM:   request.UOM,
		Price: request.Price,
	}

	id, err := uf.userRepo.InsertUserRepository(ctx, product)
	if err != nil {
		return
	}

	resp = model.AddedUserResponse{
		Id:   id,
		Name: product.Name,
	}

	// userId := 1
	// Check Health sales
	// lf.queueService.PublishData(ctx, constant.CONSUMER_PRODUCT_INSERT_RABBITMQ, userId)

	return
}
