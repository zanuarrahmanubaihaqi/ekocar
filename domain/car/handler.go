package car

import (
	"eko-car/domain/shared/context"
	Error "eko-car/domain/shared/error"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/shared/response"
	"eko-car/domain/car/constant"
	"eko-car/domain/car/feature"
	"eko-car/domain/car/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CarHandler interface {
	AddCarHandler(c *fiber.Ctx) error
	GetCarHandler(c *fiber.Ctx) error
	GetCarListsHandler(c *fiber.Ctx) error
	UpdateCarHandler(c *fiber.Ctx) error
	DeleteCarHandler(c *fiber.Ctx) error
	GetCarListsWithFilterHandler(c *fiber.Ctx) error
}

type carHandler struct {
	feature feature.CarFeature
}

func NewCarHandler(feature feature.CarFeature) CarHandler {
	return &carHandler{
		feature: feature,
	}
}

func (lh carHandler) AddCarHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	request := new(model.AddCarRequest)
	if err := c.BodyParser(request); err != nil {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	} else if request.Name == "" {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.AddCarFeature(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgInsertDataSuccess, results)
}

func (lh carHandler) GetCarHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	sid := c.Params("id")
	if sid == "" || sid == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrCarIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	id, _ := strconv.Atoi(sid)
	results, err := lh.feature.GetCarFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetCarSuccess, results)
}

func (lh carHandler) GetCarListsHandler(c *fiber.Ctx) error {

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

	resp, err := lh.feature.GetCarListsFeature(ctx, queryRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetListsDataSuccess, resp)
}

func (lh carHandler) DeleteCarHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	sid := c.Params("id")
	if sid == "" || sid == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrCarIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	id, _ := strconv.Atoi(sid)
	results, err := lh.feature.DeleteCarFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgDeleteCarSuccess, results)
}

func (lh carHandler) UpdateCarHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	sid := c.Params("id")
	if sid == "" || sid == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrCarIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	request := new(model.UpdateCarRequest)
	if err := c.BodyParser(request); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrCarIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	id, _ := strconv.Atoi(sid)
	results, err := lh.feature.UpdateCarFeature(ctx, id, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgUpdateCarSuccess, results)
}

func (lh carHandler) GetCarListsWithFilterHandler(c *fiber.Ctx) error {

	ctx := context.CreateContext()
	ctx = context.SetValueToContext(ctx, c)

	filterRequest := new(shared_model.Filter)
	if err := c.BodyParser(filterRequest); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := lh.feature.GetListsCarWithFilters(ctx, filterRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}
