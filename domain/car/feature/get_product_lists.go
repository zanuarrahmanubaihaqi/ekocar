package feature

import (
	"eko-car/domain/car/constant"
	"eko-car/domain/car/model"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/shared/helper"
	shared_model "eko-car/domain/shared/model"
	"context"
	"strings"
)

func (lf carFeature) GetCarListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (productList model.CarLists, err error) {

	// Cleaning & Set Sort query
	var (
		qsortList     []string
		qFilterList   []string
		sortby        = queryRequest.SortBy
		search        = queryRequest.Search
		totalCars int
	)

	// Sort
	sortby = strings.TrimSpace(sortby)
	if sortby != "" {
		sortby, qsortList, err = helper.SortBy(sortby)
		if err != nil {
			err = Error.New(constant.ErrGeneral, constant.ErrInvalidSortBy, err)
		}
	}

	// Search
	if search != "" {
		search, qFilterList, err = helper.FilterBy(search)
		if err != nil {
			err = Error.New(constant.ErrGeneral, constant.ErrInvalidFilterBy, err)
		}

		// Get Total Car Now
		totalCars, err = lf.carRepo.GetTotalCarWithConditionsRepository(ctx, search)
		if err != nil {
			return
		}
	} else {
		// Get Total Car Now
		totalCars, err = lf.carRepo.GetTotalCarRepository(ctx)
		if err != nil {
			return
		}
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalCars, queryRequest.Limit, queryRequest.Page)

	// Get Lists Car
	products, err := lf.carRepo.GetCarListsRepository(ctx, queryRequest.Limit, offset, sortby, search)
	if err != nil {
		return
	}

	productList = model.CarLists{
		Pagination: shared_model.Pagination{
			Limit:     queryRequest.Limit,
			TotalPage: total_page,
			TotalRows: totalCars,
			Page:      queryRequest.Page,
		},
		Car: products,
		Sort:    qsortList,
		Filter:  qFilterList,
	}

	return
}

func (lf carFeature) GetListsCarWithFilters(ctx context.Context, filter *shared_model.Filter) (productList model.CarListsByFilter, err error) {

	// Get Total Car Now
	totalCars, err := lf.carRepo.GetTotalCarWithFiltersRepository(ctx, filter)
	if err != nil {
		return
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalCars, filter.Limit, filter.Page)

	// Get Lists Car
	products, err := lf.carRepo.GetCarListsWithFiltersRepository(ctx, filter, offset)
	if err != nil {
		return
	}

	productList = model.CarListsByFilter{
		Pagination: shared_model.Pagination{
			Limit:     filter.Limit,
			TotalPage: total_page,
			TotalRows: totalCars,
			Page:      filter.Page,
		},
		Car: products,
		Filters: filter.Filters,
	}

	return
}
