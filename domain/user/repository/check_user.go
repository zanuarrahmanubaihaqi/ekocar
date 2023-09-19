package repository

import (
	"context"
	"database/sql"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
)

func (lr userRepository) CheckUserIdRepository(ctx context.Context, id int) (exist bool, err error) {
	rows, err := lr.Database.QueryContext(ctx, "SELECT COUNT(*) FROM user WHERE deleted_at IS NULL AND id = ? LIMIT 1", &id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return false, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}

		if count == 1 {
			exist = true
			break
		}
	}

	return
}
