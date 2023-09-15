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
	Car    []Car        `json:"products"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type CarListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Car    []Car        `json:"products"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}
