package repository

import (
	"context"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/user/model"
	"fmt"
	"strings"
)

func (lr userRepository) UpdateUserRepository(ctx context.Context, id int, update *model.UpdateUserRequest) (err error) {

	buildQuery := []string{}
	if update.Name != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("name = %s", update.Name))
	}
	if update.Email != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("email = %s", update.Email))
	}
	if update.NoTlp != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("no_tlp = %s", update.NoTlp))
	}
	if update.Status != 0 {
		buildQuery = append(buildQuery, fmt.Sprintf("status = %d", update.Status))
	}
	if update.Password != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("password = %s", update.Password))
	}
	if update.ValidPassword != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("valid_password = %s", update.ValidPassword))
	}
	if update.Address != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("address = %s", update.Address))
	}
	if update.NoKtp != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("no_ktp = %s", update.NoKtp))
	}
	if update.NoRekening != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("no_rekening = %s", update.NoRekening))
	}
	if update.Role != 0 {
		buildQuery = append(buildQuery, fmt.Sprintf("role = %d", update.Role))
	}

	updateQuery := strings.Join(buildQuery, ",")
	query := fmt.Sprintf("UPDATE user SET %s , updated_at = now() WHERE id = ?", updateQuery)

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
