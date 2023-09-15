package repository

import (
	"eko-car/domain/logistik/model"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/shared/query"
	"eko-car/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (lr logistikRepository) GetProductListsRepository(ctx context.Context, limit, offset int, sortby, search string) (products []model.Product, err error) {

	if sortby == "" {
		sortby = "id asc"
	}

	if search != "" {
		search = query.SearchQueryBuilder(search)
	}

	query := fmt.Sprintf("SELECT * FROM Product WHERE deleted_at IS NULL %s  ORDER BY %s LIMIT $1 OFFSET $2", search, sortby)
	logger.LogInfo(constant.QUERY, query)

	err = lr.Database.DB.SelectContext(ctx, &products, query, limit, offset)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return products, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}

func (lr logistikRepository) GetProductListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (products []model.Product, err error) {

	query, err := query.SelectStatementBuilder(model.Product{}, "product", filter)
	if err != nil {
		err = Error.New(constant.ErrDatabase, "error when create select statements", err)
		return
	}

	logger.LogInfo(constant.QUERY, query)
	if len(filter.Filters) > 0 {
		err = lr.Database.DB.SelectContext(ctx, &products, query, offset)
	} else {
		err = lr.Database.DB.SelectContext(ctx, &products, query)
	}

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return products, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}
