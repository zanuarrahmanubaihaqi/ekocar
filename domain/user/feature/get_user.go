package feature

import (
	"context"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/user/constant"
	"eko-car/domain/user/model"
	"errors"
	"strconv"
)

func (uf userFeature) GetUserFeature(ctx context.Context, id int) (response model.User, err error) {

	result, err := uf.userRepo.GetUserByIdRepository(ctx, id)
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
