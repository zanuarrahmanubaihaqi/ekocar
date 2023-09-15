package feature

import (
	"eko-car/domain/car/constant"
	"eko-car/domain/car/model"
	Error "eko-car/domain/shared/error"
	"context"
	"strconv"
)

func (lf carFeature) DeleteCarFeature(ctx context.Context, id string) (response model.DeletedCarResponse, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	err = lf.carRepo.DeleteCarRepository(ctx, idInt)
	if err != nil {
		return
	}

	response = model.DeletedCarResponse{
		Id: idInt,
	}

	return
}
