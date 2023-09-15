package feature

import (
	"eko-car/domain/user/constant"
	"eko-car/domain/user/model"
	Error "eko-car/domain/shared/error"
	"context"
	"strconv"
)

func (uf userFeature) DeleteUserFeature(ctx context.Context, id string) (response model.DeletedUserResponse, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	err = uf.userRepo.DeleteUserRepository(ctx, idInt)
	if err != nil {
		return
	}

	response = model.DeletedUserResponse{
		Id: idInt,
	}

	return
}
