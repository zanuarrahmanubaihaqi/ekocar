package feature

import (
	"context"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/car/constant"
	"eko-car/domain/car/model"
	"errors"
	"strconv"
)

func (uf carFeature) UpdateCarFeature(ctx context.Context, id int, request *model.UpdateCarRequest) (response model.Car, err error) {

	// Check Car Id
	exist, err := uf.carRepo.GetCarByIdRepository(ctx, id)
	if err != nil {
		return
	} else if exist.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrCarIdNotFound, errors.New("not found"))
		return
	}

	// Update Car
	err = uf.carRepo.UpdateCarRepository(ctx, id, request)
	if err != nil {
		return
	}

	if exist.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrCarIdNotFound, errors.New(strconv.Itoa(exist.Id)))
		return
	}

	response = exist
	return

}
