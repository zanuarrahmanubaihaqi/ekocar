package feature

import (
	"context"
	"eko-car/domain/car/model"
)

func (uf carFeature) DeleteCarFeature(ctx context.Context, id int) (response model.DeletedCarResponse, err error) {

	err = uf.carRepo.DeleteCarRepository(ctx, id)
	if err != nil {
		return
	}

	response = model.DeletedCarResponse{
		Id: id,
	}

	return
}
