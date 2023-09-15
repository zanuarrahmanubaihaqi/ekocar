package user

import (
	"eko-car/domain/user/constant"
	"eko-car/domain/user/feature"
	"eko-car/domain/user/model"
	"eko-car/domain/shared/context"
	Error "eko-car/domain/shared/error"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/shared/response"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	AddUserHandler(c *fiber.Ctx) error
	GetUserHandler(c *fiber.Ctx) error
	GetUserListsHandler(c *fiber.Ctx) error
	UpdateUserHandler(c *fiber.Ctx) error
	DeleteUserHandler(c *fiber.Ctx) error
	BulkCounterHandler(c *fiber.Ctx) error
	GetUserListsWithFilterHandler(c *fiber.Ctx) error
}

type userHandler struct {
	feature feature.UserFeature
}

func NewUserHandler(feature feature.UserFeature) UserHandler {
	return &userHandler{
		feature: feature,
	}
}

func (lh userHandler) AddUserHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	request := new(model.AddUserRequest)
	if err := c.BodyParser(request); err != nil {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	} else if request.Name == "" || request.SKU == "" || request.UOM == "" {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.AddUserFeature(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgInsertDataSuccess, results)
}

func (lh userHandler) GetUserHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrUserIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.GetUserFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetUserSuccess, results)
}

func (lh userHandler) GetUserListsHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	page, err := strconv.Atoi(strings.TrimSpace(c.Query(constant.PAGE)))
	if err != nil || page == 0 {
		page = constant.DefaultPage
	}

	limit, err := strconv.Atoi(strings.TrimSpace(c.Query(constant.LIMIT)))
	if err != nil || limit == 0 {
		limit = constant.DefaultLimitPerPage
	}

	sortBy := strings.TrimSpace(c.Query(constant.SORT_BY))
	search := strings.TrimSpace(c.Query(constant.SEARCH))

	queryRequest := shared_model.QueryRequest{
		Page:   page,
		Limit:  limit,
		SortBy: sortBy,
		Search: search,
	}

	resp, err := lh.feature.GetUserListsFeature(ctx, queryRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetListsDataSuccess, resp)
}

func (lh userHandler) DeleteUserHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrUserIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.DeleteUserFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgDeleteUserSuccess, results)
}

func (lh userHandler) UpdateUserHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrUserIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	request := new(model.UpdateUserRequest)
	if err := c.BodyParser(request); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrUserIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.UpdateUserFeature(ctx, id, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgUpdateUserSuccess, results)
}

func (lh userHandler) BulkCounterHandler(c *fiber.Ctx) error {

	ctx := context.CreateContext()
	ctx = context.SetValueToContext(ctx, c)

	err := lh.feature.BulkCounterFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, "request being processed")
}

func (lh userHandler) GetUserListsWithFilterHandler(c *fiber.Ctx) error {

	ctx := context.CreateContext()
	ctx = context.SetValueToContext(ctx, c)

	filterRequest := new(shared_model.Filter)
	if err := c.BodyParser(filterRequest); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := lh.feature.GetListsUserWithFilters(ctx, filterRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}
