package repository

import (
	"eko-car/domain/car/model"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	"context"
	"fmt"
	"strings"
)

func (lr carRepository) UpdateCarRepository(ctx context.Context, id int, update *model.UpdateCarRequest) (err error) {

	buildQuery := []string{}
	if update.Name != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("name = '%s'", update.Name))
	}
	if update.SKU != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("sku ='%s'", update.SKU))
	}
	if update.Price != 0 {
		buildQuery = append(buildQuery, fmt.Sprintf("price = %d", update.Price))
	}
	if update.UOM != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("uom = '%s'", update.UOM))
	}

	updateQuery := strings.Join(buildQuery, ",")
	query := fmt.Sprintf("UPDATE product SET %s , updated_at = now() WHERE id = $1", updateQuery)

	tx := lr.Database.DB.MustBegin()
	_, err = tx.QueryContext(ctx, query, &id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			tx.Rollback()
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenCommitDB, err)
			tx.Rollback()
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}
