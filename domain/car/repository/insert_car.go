package repository

import (
	"context"
	"eko-car/domain/car/model"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	"time"
)

var now = time.Now()

func (lr *carRepository) InsertCarRepository(ctx context.Context, request model.Car) (id int, err error) {
	id = 0
	tx := lr.Database.DB.MustBegin()
	query := `
		INSERT INTO 
			car (
					merk,
					jeins,
					type,
					tahun_pembuatan,
					image,
					harga,
					lokasi,
					komisi,
					deskripsi,
					created_at
				) VALUES (?,?,?,?,?,?,?,?,?,?)`
	stmt, err := lr.Database.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = stmt.ExecContext(ctx, &request.Merk, &request.Jenis, &request.Type, &request.TahunPembuatan, &request.Image, &request.Harga, &request.Lokasi, &request.Komisi, &request.Deskripsi, now.Format("2006-01-2 15:04:05"))

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
