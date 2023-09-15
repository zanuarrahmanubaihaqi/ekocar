package feature

import (
	"eko-car/domain/user/constant"
	"eko-car/domain/user/model"
	Error "eko-car/domain/shared/error"
	"context"
	"errors"
	"strconv"
)

func (uf userFeature) GetUserFeature(ctx context.Context, id string) (response model.User, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := uf.userRepo.GetUserByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	if result.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrUserIdNotFound, errors.New(strconv.Itoa(result.Id)))
		return
	}

	response = result

	return
}
