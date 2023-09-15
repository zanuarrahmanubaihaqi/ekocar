package feature

import (
	"context"
)

func (uf userFeature) BulkCounterFeature(ctx context.Context) (err error) {
	var (
		size = 1000
	)

	go uf.userRepo.BulkInsertCounterRepository(ctx, size)

	return
}
