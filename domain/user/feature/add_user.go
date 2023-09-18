package feature

import (
	"context"
	"eko-car/domain/user/model"

	"github.com/thanhpk/randstr"
)

func (uf userFeature) AddUserFeature(ctx context.Context, request *model.AddUserRequest) (resp model.AddedUserResponse, err error) {

	// Added Bussiness logic here

	randStr := randstr.String(20)
	parameter := model.User{
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
		UniqueCode:    randStr,
	}

	id, err := uf.userRepo.InsertUserRepository(ctx, parameter)
	if err != nil {
		return
	}

	resp = model.AddedUserResponse{
		Id:   id,
		Name: parameter.Name,
	}

	return
}
