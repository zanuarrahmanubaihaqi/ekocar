package logistik

import (
	"eko-car/domain/logistik/constant"
	"eko-car/domain/logistik/feature"
	"eko-car/domain/logistik/model"
	"eko-car/domain/shared/context"
	Error "eko-car/domain/shared/error"
	shared_model "eko-car/domain/shared/model"
	"eko-car/domain/shared/response"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type LogistikHandler interface {
	AddProductHandler(c *fiber.Ctx) error
	GetProductHandler(c *fiber.Ctx) error
	GetProductListsHandler(c *fiber.Ctx) error
	UpdateProductHandler(c *fiber.Ctx) error
	DeleteProductHandler(c *fiber.Ctx) error
	BulkCounterHandler(c *fiber.Ctx) error
	GetProductListsWithFilterHandler(c *fiber.Ctx) error
}

type logistikHandler struct {
	feature feature.LogistikFeature
}

func NewLogistikHandler(feature feature.LogistikFeature) LogistikHandler {
	return &logistikHandler{
		feature: feature,
	}
}

func (lh logistikHandler) AddProductHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	request := new(model.AddProductRequest)
	if err := c.BodyParser(request); err != nil {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	} else if request.Name == "" || request.SKU == "" || request.UOM == "" {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.AddProductFeature(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgInsertDataSuccess, results)
}

func (lh logistikHandler) GetProductHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrProductIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.GetProductFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetProductSuccess, results)
}

func (lh logistikHandler) GetProductListsHandler(c *fiber.Ctx) error {

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

	resp, err := lh.feature.GetProductListsFeature(ctx, queryRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetListsDataSuccess, resp)
}

func (lh logistikHandler) DeleteProductHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrProductIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.DeleteProductFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgDeleteProductSuccess, results)
}

func (lh logistikHandler) UpdateProductHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrProductIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	request := new(model.UpdateProductRequest)
	if err := c.BodyParser(request); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrProductIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.UpdateProductFeature(ctx, id, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgUpdateProductSuccess, results)
}

func (lh logistikHandler) BulkCounterHandler(c *fiber.Ctx) error {

	ctx := context.CreateContext()
	ctx = context.SetValueToContext(ctx, c)

	err := lh.feature.BulkCounterFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, "request being processed")
}

func (lh logistikHandler) GetProductListsWithFilterHandler(c *fiber.Ctx) error {

	ctx := context.CreateContext()
	ctx = context.SetValueToContext(ctx, c)

	filterRequest := new(shared_model.Filter)
	if err := c.BodyParser(filterRequest); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := lh.feature.GetListsProductWithFilters(ctx, filterRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}
