package feature

import (
	"context"
	"eko-car/domain/car/model"
)

func (uf carFeature) AddCarFeature(ctx context.Context, request *model.AddCarRequest) (resp model.AddedCarResponse, err error) {

	// Added Bussiness logic here

	parameter := model.Car{
		Name:          request.Name,
		Email:         request.Email,
		NoTlp:         request.NoTlp,
		Password:      request.Password,
		ValidPassword: request.ValidPassword,
		Address:       request.Address,
		NoRekening:    request.NoRekening,
		NoKtp:         request.NoKtp,
		NoNpwp:        request.NoNpwp,
		ImageKtp:      request.ImageKtp,
		Role:          request.Role,
		Status:        1,
		UniqueCode:    "",
	}

	id, err := uf.carRepo.InsertCarRepository(ctx, parameter)
	if err != nil {
		return
	}

	resp = model.AddedCarResponse{
		Id:   id,
		Name: parameter.Name,
	}

	return
}
