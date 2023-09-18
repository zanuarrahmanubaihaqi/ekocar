package feature

import (
	"context"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/user/model"
	repository "eko-car/domain/user/repository"
)

type UserFeature interface {
	AddUserFeature(ctx context.Context, request *model.AddUserRequest) (response model.AddedUserResponse, err error)
	GetUserFeature(ctx context.Context, id int) (response model.User, err error)
	GetUserListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (productList model.UserLists, err error)
	DeleteUserFeature(ctx context.Context, id int) (response model.DeletedUserResponse, err error)
	UpdateUserFeature(ctx context.Context, id int, request *model.UpdateUserRequest) (response model.User, err error)
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
