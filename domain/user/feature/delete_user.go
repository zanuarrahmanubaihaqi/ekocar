package feature

import (
	"context"
	"eko-car/domain/user/model"
)

func (uf userFeature) DeleteUserFeature(ctx context.Context, id int) (response model.DeletedUserResponse, err error) {

	err = uf.userRepo.DeleteUserRepository(ctx, id)
	if err != nil {
		return
	}

	response = model.DeletedUserResponse{
		Id: id,
	}

	return
}
