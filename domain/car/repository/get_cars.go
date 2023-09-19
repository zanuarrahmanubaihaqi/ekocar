package repository

import (
	"context"
	"database/sql"
	"eko-car/domain/car/model"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	sharedModel "eko-car/domain/shared/model"
	"eko-car/domain/shared/query"
	"eko-car/infrastructure/logger"
	"fmt"
)

func (lr carRepository) GetCarListsRepository(ctx context.Context, limit, offset int, sortby, search string, queryRequest sharedModel.QueryRequest) (products []model.Car, err error) {

	if sortby == "" {
		sortby = "id asc"
	}

	if search != "" {
		search = query.SearchQueryBuilder(search)
	}

	squery := ""
	if queryRequest.Merk != "" {
		squery += fmt.Sprintf(` AND merk LIKE %s`, "'%"+queryRequest.Merk+"%'")
	}

	query := fmt.Sprintf(`
		SELECT 
			id,
			merk,
			jeins,
			type,
			tahun_pembuatan,
			image,
			harga,
			lokasi,
			komisi,
			deskripsi
		FROM 
			car 
		WHERE 
			deleted_at %s IS NULL %s  ORDER BY %s LIMIT ? OFFSET ?`, squery, search, sortby)
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

func (lr carRepository) GetCarListsWithFiltersRepository(ctx context.Context, filter *sharedModel.Filter, offset int) (products []model.Car, err error) {

	query, err := query.SelectStatementBuilder(model.Car{}, "car", filter)
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
