package model

import "eko-car/domain/shared/model"

type AddedCarResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DeletedCarResponse struct {
	Id int `json:"id"`
}

type CarLists struct {
	Pagination model.Pagination `json:"pagination"`
	Car        []Car            `json:"products"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type CarListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Car        []Car            `json:"products"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}

type Car struct {
	Id             int     `json:"id" query:"id" db:"id" form:"id"`
	Merk           string  `json:"merk" query:"merk" db:"merk" form:"merk"`
	Jenis          string  `json:"jeins" query:"jeins" db:"jeins" form:"jeins"`
	Type           string  `json:"type" query:"type" db:"type" form:"type"`
	TahunPembuatan string  `json:"tahun_pembuatan" query:"tahun_pembuatan" db:"tahun_pembuatan" form:"tahun_pembuatan"`
	Image          string  `json:"image" query:"image" db:"image" form:"image"`
	Harga          float64 `json:"harga" query:"harga" db:"harga" form:"harga"`
	Lokasi         string  `json:"lokasi" query:"lokasi" db:"lokasi" form:"lokasi"`
	Komisi         float64 `json:"komisi" query:"komisi" db:"komisi" form:"komisi"`
	Deskripsi      string  `json:"deskripsi" query:"deskripsi" db:"deskripsi" form:"deskripsi"`
}
