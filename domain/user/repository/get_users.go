package repository

import (
	"context"
	"database/sql"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	sharedModel "eko-car/domain/shared/model"
	"eko-car/domain/shared/query"
	"eko-car/domain/user/model"
	"eko-car/infrastructure/logger"
	"fmt"
)

func (lr userRepository) GetUserListsRepository(ctx context.Context, limit, offset int, sortby, search string, queryRequest sharedModel.QueryRequest) (products []model.User, err error) {

	if sortby == "" {
		sortby = "id asc"
	}

	if search != "" {
		search = query.SearchQueryBuilder(search)
	}

	squery := ""
	if queryRequest.Name != "" {
		squery += fmt.Sprintf(` AND name LIKE %s`, "'%"+queryRequest.Name+"%'")
	}

	query := fmt.Sprintf(`
		SELECT 
			name,
			email,
			no_tlp,
			status,
			password,
			valid_password,
			unique_code,
			address,
			no_ktp,
			image_ktp,
			no_npwp,
			no_rekening,
			role
		FROM 
			user 
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

func (lr userRepository) GetUserListsWithFiltersRepository(ctx context.Context, filter *sharedModel.Filter, offset int) (products []model.User, err error) {

	query, err := query.SelectStatementBuilder(model.User{}, "user", filter)
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
