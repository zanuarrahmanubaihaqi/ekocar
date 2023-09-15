package feature

import (
	"eko-car/domain/user/constant"
	"eko-car/domain/user/model"
	Error "eko-car/domain/shared/error"
	"context"
	"errors"
	"strconv"
	"strings"
)

func (uf userFeature) UpdateUserFeature(ctx context.Context, id string, request *model.UpdateUserRequest) (response model.User, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	// Check User Id
	exist, err := uf.userRepo.CheckUserIdRepository(ctx, idInt)
	if err != nil {
		return
	} else if !exist {
		err = Error.New(constant.ErrGeneral, constant.ErrUserIdNotFound, errors.New(id))
		return
	}

	// Check User SKU
	if strings.TrimSpace(request.SKU) != "" {
		exist, err = uf.userRepo.CheckUserSKURepository(ctx, request.SKU)
		if err != nil {
			return
		} else if exist {
			err = Error.New(constant.ErrGeneral, constant.ErrSKUAlreadyExist, errors.New(request.SKU))
			return
		}
	}

	// Update User
	err = uf.userRepo.UpdateUserRepository(ctx, idInt, request)
	if err != nil {
		return
	}

	// Get New User
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
