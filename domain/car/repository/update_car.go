package repository

import (
	"context"
	"eko-car/domain/car/model"
	"eko-car/domain/shared/constant"
	Error "eko-car/domain/shared/error"
	"fmt"
	"strings"
)

func (lr carRepository) UpdateCarRepository(ctx context.Context, id int, update *model.UpdateCarRequest) (err error) {

	buildQuery := []string{}
	if update.Merk != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("merk = %s", update.Merk))
	}
	if update.Jenis != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("jenis = %s", update.Jenis))
	}
	if update.Type != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("type = %s", update.Type))
	}
	if update.TahunPembuatan != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("tahun_pembuatan = %d", update.TahunPembuatan))
	}
	if update.Image != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("image = %s", update.Image))
	}
	if update.Harga != 0 {
		buildQuery = append(buildQuery, fmt.Sprintf("harga = %f", update.Harga))
	}
	if update.Lokasi != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("lokasi = %s", update.Lokasi))
	}
	if update.Komisi != 0 {
		buildQuery = append(buildQuery, fmt.Sprintf("komisi = %f", update.Komisi))
	}
	if update.Deskripsi != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("deskripsi = %s", update.Deskripsi))
	}

	updateQuery := strings.Join(buildQuery, ",")
	query := fmt.Sprintf("UPDATE car SET %s , updated_at = now() WHERE id = ?", updateQuery)

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
