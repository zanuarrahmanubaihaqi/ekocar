package feature

import (
	"eko-car/domain/car/constant"
	"eko-car/domain/car/model"
	Error "eko-car/domain/shared/error"
	"context"
	"errors"
	"strconv"
	"strings"
)

func (lf carFeature) UpdateCarFeature(ctx context.Context, id string, request *model.UpdateCarRequest) (response model.Car, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	// Check Car Id
	exist, err := lf.carRepo.CheckCarIdRepository(ctx, idInt)
	if err != nil {
		return
	} else if !exist {
		err = Error.New(constant.ErrGeneral, constant.ErrCarIdNotFound, errors.New(id))
		return
	}

	// Check Car SKU
	if strings.TrimSpace(request.SKU) != "" {
		exist, err = lf.carRepo.CheckCarSKURepository(ctx, request.SKU)
		if err != nil {
			return
		} else if exist {
			err = Error.New(constant.ErrGeneral, constant.ErrSKUAlreadyExist, errors.New(request.SKU))
			return
		}
	}

	// Update Car
	err = lf.carRepo.UpdateCarRepository(ctx, idInt, request)
	if err != nil {
		return
	}

	// Get New Car
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
