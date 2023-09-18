package feature

import (
	"context"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/shared/helper"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/user/constant"
	"eko-car/domain/user/model"
	"strings"
)

func (uf userFeature) GetUserListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (productList model.UserLists, err error) {

	// Cleaning & Set Sort query
	var (
		qsortList   []string
		qFilterList []string
		sortby      = queryRequest.SortBy
		search      = queryRequest.Search
		totalUsers  int
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

		// Get Total User Now
		totalUsers, err = uf.userRepo.GetTotalUserWithConditionsRepository(ctx, search)
		if err != nil {
			return
		}
	} else {
		// Get Total User Now
		totalUsers, err = uf.userRepo.GetTotalUserRepository(ctx)
		if err != nil {
			return
		}
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalUsers, queryRequest.Limit, queryRequest.Page)

	// Get Lists User
	products, err := uf.userRepo.GetUserListsRepository(ctx, queryRequest.Limit, offset, sortby, search, queryRequest)
	if err != nil {
		return
	}

	productList = model.UserLists{
		Pagination: shared_model.Pagination{
			Limit:     queryRequest.Limit,
			TotalPage: total_page,
			TotalRows: totalUsers,
			Page:      queryRequest.Page,
		},
		User:   products,
		Sort:   qsortList,
		Filter: qFilterList,
	}

	return
}

func (uf userFeature) GetListsUserWithFilters(ctx context.Context, filter *shared_model.Filter) (productList model.UserListsByFilter, err error) {

	// Get Total User Now
	totalUsers, err := uf.userRepo.GetTotalUserWithFiltersRepository(ctx, filter)
	if err != nil {
		return
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalUsers, filter.Limit, filter.Page)

	// Get Lists User
	products, err := uf.userRepo.GetUserListsWithFiltersRepository(ctx, filter, offset)
	if err != nil {
		return
	}

	productList = model.UserListsByFilter{
		Pagination: shared_model.Pagination{
			Limit:     filter.Limit,
			TotalPage: total_page,
			TotalRows: totalUsers,
			Page:      filter.Page,
		},
		User:    products,
		Filters: filter.Filters,
	}

	return
}
