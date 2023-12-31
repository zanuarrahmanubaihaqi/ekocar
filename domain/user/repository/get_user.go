package repository

import (
	"context"
	"database/sql"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/shared/query"
	"eko-car/domain/user/model"
	"eko-car/infrastructure/logger"
	"fmt"
)

func (lr userRepository) GetUserByIdRepository(ctx context.Context, id int) (product model.User, err error) {

	query := `
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
			id = ? 
			AND deleted_at IS NULL 
		LIMIT 1`
	logger.LogInfo(constant.QUERY, query)

	rows, err := lr.Database.Queryx(query, &id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return product, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err = rows.StructScan(&product)
		if err != nil {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr userRepository) GetTotalUserRepository(ctx context.Context) (count int, err error) {

	query := "SELECT COUNT(*) FROM user WHERE deleted_at IS NULL"
	rows, err := lr.Database.Query(query)
	logger.LogInfo(constant.QUERY, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return 0, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr userRepository) GetTotalUserWithConditionsRepository(ctx context.Context, conditions string) (count int, err error) {

	if conditions != "" {
		conditions = query.SearchQueryBuilder(conditions)
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM product WHERE deleted_at IS NULL %s", conditions)
	logger.LogInfo(constant.QUERY, query)

	rows, err := lr.Database.Query(query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return 0, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr userRepository) GetTotalUserWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error) {

	var (
		conditions string
	)

	if filter != nil {
		conditions = query.ConditionsBuilder(filter)
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM product WHERE deleted_at IS NULL %s", "")
	if len(filter.Filters) > 0 {
		query = fmt.Sprintf("SELECT COUNT(*) FROM product WHERE deleted_at IS NULL AND %s", conditions)
	}

	logger.LogInfo(constant.QUERY, query)
	rows, err := lr.Database.Query(query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return 0, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}
