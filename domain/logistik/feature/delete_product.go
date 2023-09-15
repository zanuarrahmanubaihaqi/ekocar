package feature

import (
	"eko-car/domain/logistik/constant"
	"eko-car/domain/logistik/model"
	Error "eko-car/domain/shared/error"
	"context"
	"strconv"
)

func (lf logistikFeature) DeleteProductFeature(ctx context.Context, id string) (response model.DeletedProductResponse, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	err = lf.logistikRepo.DeleteProductRepository(ctx, idInt)
	if err != nil {
		return
	}

	response = model.DeletedProductResponse{
		Id: idInt,
	}

	return
}
