package repository

import (
	"context"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/user/model"
	"eko-car/infrastructure/database"
)

type UserRepository interface {
	InsertUserRepository(ctx context.Context, product model.User) (id int, err error)
	GetUserByIdRepository(ctx context.Context, id int) (product model.User, err error)
	GetUserListsRepository(ctx context.Context, limit, offset int, sortby, search string, queryRequest shared_model.QueryRequest) (products []model.User, err error)
	GetTotalUserRepository(ctx context.Context) (count int, err error)
	DeleteUserRepository(ctx context.Context, id int) (err error)
	UpdateUserRepository(ctx context.Context, id int, update *model.UpdateUserRequest) (err error)
	CheckUserIdRepository(ctx context.Context, id int) (exist bool, err error)
	GetTotalUserWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	GetTotalUserWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	GetUserListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (products []model.User, err error)
}

type userRepository struct {
	Database *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{
		Database: db,
	}
}
