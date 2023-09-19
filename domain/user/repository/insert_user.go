package repository

import (
	"context"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/user/model"
	"time"
)

var now = time.Now()

func (lr *userRepository) InsertUserRepository(ctx context.Context, request model.User) (id int, err error) {
	id = 0
	tx := lr.Database.DB.MustBegin()
	query := `
		INSERT INTO 
			user (
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
					role,
					created_at
				) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	stmt, err := lr.Database.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = stmt.ExecContext(ctx, &request.Name, &request.Email, &request.NoTlp, &request.Status, &request.Password, &request.ValidPassword, &request.UniqueCode, &request.Address, &request.NoKtp, &request.ImageKtp, &request.NoNpwp, &request.NoRekening, &request.Role, now.Format("2006-01-2 15:04:05"))

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			tx.Rollback()
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}
