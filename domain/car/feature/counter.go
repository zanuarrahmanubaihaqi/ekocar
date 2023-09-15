package feature

import (
	"context"
)

func (lf carFeature) BulkCounterFeature(ctx context.Context) (err error) {
	var (
		size = 1000
	)

	go lf.carRepo.BulkInsertCounterRepository(ctx, size)

	return
}
