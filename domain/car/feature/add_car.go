package feature

import (
	"context"
	"eko-car/domain/car/model"
)

func (uf carFeature) AddCarFeature(ctx context.Context, request *model.AddCarRequest) (resp model.AddedCarResponse, err error) {

	// Added Bussiness logic here

	parameter := model.Car{
		Merk:           request.Merk,
		Jenis:          request.Jenis,
		Type:           request.Type,
		TahunPembuatan: request.TahunPembuatan,
		Image:          request.Image,
		Harga:          request.Harga,
		Lokasi:         request.Lokasi,
		Komisi:         request.Komisi,
	}

	id, err := uf.carRepo.InsertCarRepository(ctx, parameter)
	if err != nil {
		return
	}

	resp = model.AddedCarResponse{
		Id:   id,
		Name: parameter.Merk,
	}

	return
}
