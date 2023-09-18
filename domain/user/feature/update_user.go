package feature

import (
	"context"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/user/constant"
	"eko-car/domain/user/model"
	"errors"
	"strconv"
)

func (uf userFeature) UpdateUserFeature(ctx context.Context, id int, request *model.UpdateUserRequest) (response model.User, err error) {

	// Check User Id
	exist, err := uf.userRepo.GetUserByIdRepository(ctx, id)
	if err != nil {
		return
	} else if exist.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrUserIdNotFound, errors.New("not found"))
		return
	}

	// Update User
	err = uf.userRepo.UpdateUserRepository(ctx, id, request)
	if err != nil {
		return
	}

	if exist.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrUserIdNotFound, errors.New(strconv.Itoa(exist.Id)))
		return
	}

	response = exist
	return

}
