package feature

import (
	"eko-car/domain/car/constant"
	"eko-car/domain/car/model"
	Error "eko-car/domain/shared/error"
	"context"
	"errors"
	"strconv"
)

func (lf carFeature) GetCarFeature(ctx context.Context, id string) (response model.Car, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := lf.carRepo.GetCarByIdRepository(ctx, idInt)
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
