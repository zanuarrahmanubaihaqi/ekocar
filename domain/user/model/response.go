package model

import "eko-car/domain/shared/model"

type AddedUserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DeletedUserResponse struct {
	Id int `json:"id"`
}

type UserLists struct {
	Pagination model.Pagination `json:"pagination"`
	User       []User           `json:"products"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type UserListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	User       []User           `json:"products"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}

type User struct {
	Id            int    `json:"id" query:"id" db:"id" form:"id"`
	Name          string `json:"name" query:"name" db:"name" form:"name"`
	Email         string `json:"email" query:"email" db:"email" form:"email"`
	NoTlp         string `json:"no_tlp" query:"no_tlp" db:"no_tlp" form:"no_tlp"`
	Status        int    `json:"status" query:"status" db:"status" form:"status"`
	Password      string `json:"password" query:"password" db:"password" form:"password"`
	ValidPassword string `json:"valid_password" query:"valid_password" db:"valid_password" form:"valid_password"`
	UniqueCode    string `json:"unique_code" query:"unique_code" db:"unique_code" form:"unique_code"`
	Address       string `json:"address" query:"address" db:"address" form:"address"`
	NoKtp         string `json:"no_ktp" query:"no_ktp" db:"no_ktp" form:"no_ktp"`
	ImageKtp      string `json:"image_ktp" query:"image_ktp" db:"image_ktp" form:"image_ktp"`
	NoNpwp        string `json:"no_npwp" query:"no_npwp" db:"no_npwp" form:"no_npwp"`
	NoRekening    string `json:"no_rekening" query:"no_rekening" db:"no_rekening" form:"no_rekening"`
	Role          int    `json:"role" query:"role" db:"role" form:"role"`
}
