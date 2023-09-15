package feature

import (
	"context"
	"eko-car/domain/user/model"
	repository "eko-car/domain/user/repository"
	shared_model "eko-car/domain/shared/model"
)

type UserFeature interface {
	AddUserFeature(ctx context.Context, request *model.AddUserRequest) (response model.AddedUserResponse, err error)
	GetUserFeature(ctx context.Context, id string) (response model.User, err error)
	GetUserListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (productList model.UserLists, err error)
	DeleteUserFeature(ctx context.Context, id string) (response model.DeletedUserResponse, err error)
	UpdateUserFeature(ctx context.Context, id string, request *model.UpdateUserRequest) (response model.User, err error)
	BulkCounterFeature(ctx context.Context) (err error)
	GetListsUserWithFilters(ctx context.Context, filter *shared_model.Filter) (productList model.UserListsByFilter, err error)
}

type userFeature struct {
	userRepo repository.UserRepository
}

func NewUserFeature(userRepo repository.UserRepository) UserFeature {
	return &userFeature{
		userRepo: userRepo,
	}
}
