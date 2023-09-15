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
	User    []User        `json:"products"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type UserListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	User    []User        `json:"products"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}
