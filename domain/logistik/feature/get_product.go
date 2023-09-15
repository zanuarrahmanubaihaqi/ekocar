package feature

import (
	"eko-car/domain/logistik/constant"
	"eko-car/domain/logistik/model"
	Error "eko-car/domain/shared/error"
	"context"
	"errors"
	"strconv"
)

func (lf logistikFeature) GetProductFeature(ctx context.Context, id string) (response model.Product, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := lf.logistikRepo.GetProductByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	if result.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrProductIdNotFound, errors.New(strconv.Itoa(result.Id)))
		return
	}

	response = result

	return
}
