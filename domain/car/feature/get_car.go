package feature

import (
	"context"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/car/constant"
	"eko-car/domain/car/model"
	"errors"
	"strconv"
)

func (uf carFeature) GetCarFeature(ctx context.Context, id int) (response model.Car, err error) {

	result, err := uf.carRepo.GetCarByIdRepository(ctx, id)
	if err != nil {
		return
	}

	if result.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrCarIdNotFound, errors.New(strconv.Itoa(result.Id)))
		return
	}

	response = result

	return
}
