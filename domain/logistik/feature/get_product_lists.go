package feature

import (
	"eko-car/domain/logistik/constant"
	"eko-car/domain/logistik/model"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/shared/helper"
	shared_model "eko-car/domain/shared/model"
	"context"
	"strings"
)

func (lf logistikFeature) GetProductListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (productList model.ProductLists, err error) {

	// Cleaning & Set Sort query
	var (
		qsortList     []string
		qFilterList   []string
		sortby        = queryRequest.SortBy
		search        = queryRequest.Search
		totalProducts int
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

		// Get Total Product Now
		totalProducts, err = lf.logistikRepo.GetTotalProductWithConditionsRepository(ctx, search)
		if err != nil {
			return
		}
	} else {
		// Get Total Product Now
		totalProducts, err = lf.logistikRepo.GetTotalProductRepository(ctx)
		if err != nil {
			return
		}
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalProducts, queryRequest.Limit, queryRequest.Page)

	// Get Lists Product
	products, err := lf.logistikRepo.GetProductListsRepository(ctx, queryRequest.Limit, offset, sortby, search)
	if err != nil {
		return
	}

	productList = model.ProductLists{
		Pagination: shared_model.Pagination{
			Limit:     queryRequest.Limit,
			TotalPage: total_page,
			TotalRows: totalProducts,
			Page:      queryRequest.Page,
		},
		Product: products,
		Sort:    qsortList,
		Filter:  qFilterList,
	}

	return
}

func (lf logistikFeature) GetListsProductWithFilters(ctx context.Context, filter *shared_model.Filter) (productList model.ProductListsByFilter, err error) {

	// Get Total Product Now
	totalProducts, err := lf.logistikRepo.GetTotalProductWithFiltersRepository(ctx, filter)
	if err != nil {
		return
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalProducts, filter.Limit, filter.Page)

	// Get Lists Product
	products, err := lf.logistikRepo.GetProductListsWithFiltersRepository(ctx, filter, offset)
	if err != nil {
		return
	}

	productList = model.ProductListsByFilter{
		Pagination: shared_model.Pagination{
			Limit:     filter.Limit,
			TotalPage: total_page,
			TotalRows: totalProducts,
			Page:      filter.Page,
		},
		Product: products,
		Filters: filter.Filters,
	}

	return
}
