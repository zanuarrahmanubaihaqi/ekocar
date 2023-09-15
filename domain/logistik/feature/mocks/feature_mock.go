// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/logistik/feature/feature.go

// Package mock_feature is a generated GoMock package.
package mock_feature

import (
	model "eko-car/domain/logistik/model"
	model0 "eko-car/domain/shared/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogistikFeature is a mock of LogistikFeature interface.
type MockLogistikFeature struct {
	ctrl     *gomock.Controller
	recorder *MockLogistikFeatureMockRecorder
}

// MockLogistikFeatureMockRecorder is the mock recorder for MockLogistikFeature.
type MockLogistikFeatureMockRecorder struct {
	mock *MockLogistikFeature
}

// NewMockLogistikFeature creates a new mock instance.
func NewMockLogistikFeature(ctrl *gomock.Controller) *MockLogistikFeature {
	mock := &MockLogistikFeature{ctrl: ctrl}
	mock.recorder = &MockLogistikFeatureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogistikFeature) EXPECT() *MockLogistikFeatureMockRecorder {
	return m.recorder
}

// AddProductFeature mocks base method.
func (m *MockLogistikFeature) AddProductFeature(ctx context.Context, request *model.AddProductRequest) (model.AddedProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProductFeature", ctx, request)
	ret0, _ := ret[0].(model.AddedProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProductFeature indicates an expected call of AddProductFeature.
func (mr *MockLogistikFeatureMockRecorder) AddProductFeature(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProductFeature", reflect.TypeOf((*MockLogistikFeature)(nil).AddProductFeature), ctx, request)
}

// BulkCounterFeature mocks base method.
func (m *MockLogistikFeature) BulkCounterFeature(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCounterFeature", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCounterFeature indicates an expected call of BulkCounterFeature.
func (mr *MockLogistikFeatureMockRecorder) BulkCounterFeature(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCounterFeature", reflect.TypeOf((*MockLogistikFeature)(nil).BulkCounterFeature), ctx)
}

// DeleteProductFeature mocks base method.
func (m *MockLogistikFeature) DeleteProductFeature(ctx context.Context, id string) (model.DeletedProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductFeature", ctx, id)
	ret0, _ := ret[0].(model.DeletedProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProductFeature indicates an expected call of DeleteProductFeature.
func (mr *MockLogistikFeatureMockRecorder) DeleteProductFeature(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductFeature", reflect.TypeOf((*MockLogistikFeature)(nil).DeleteProductFeature), ctx, id)
}

// GetListsProductWithFilters mocks base method.
func (m *MockLogistikFeature) GetListsProductWithFilters(ctx context.Context, filter *model0.Filter) (model.ProductListsByFilter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListsProductWithFilters", ctx, filter)
	ret0, _ := ret[0].(model.ProductListsByFilter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListsProductWithFilters indicates an expected call of GetListsProductWithFilters.
func (mr *MockLogistikFeatureMockRecorder) GetListsProductWithFilters(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListsProductWithFilters", reflect.TypeOf((*MockLogistikFeature)(nil).GetListsProductWithFilters), ctx, filter)
}

// GetProductFeature mocks base method.
func (m *MockLogistikFeature) GetProductFeature(ctx context.Context, id string) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductFeature", ctx, id)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductFeature indicates an expected call of GetProductFeature.
func (mr *MockLogistikFeatureMockRecorder) GetProductFeature(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductFeature", reflect.TypeOf((*MockLogistikFeature)(nil).GetProductFeature), ctx, id)
}

// GetProductListsFeature mocks base method.
func (m *MockLogistikFeature) GetProductListsFeature(ctx context.Context, queryRequest model0.QueryRequest) (model.ProductLists, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductListsFeature", ctx, queryRequest)
	ret0, _ := ret[0].(model.ProductLists)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductListsFeature indicates an expected call of GetProductListsFeature.
func (mr *MockLogistikFeatureMockRecorder) GetProductListsFeature(ctx, queryRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductListsFeature", reflect.TypeOf((*MockLogistikFeature)(nil).GetProductListsFeature), ctx, queryRequest)
}

// UpdateProductFeature mocks base method.
func (m *MockLogistikFeature) UpdateProductFeature(ctx context.Context, id string, request *model.UpdateProductRequest) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductFeature", ctx, id, request)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductFeature indicates an expected call of UpdateProductFeature.
func (mr *MockLogistikFeatureMockRecorder) UpdateProductFeature(ctx, id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductFeature", reflect.TypeOf((*MockLogistikFeature)(nil).UpdateProductFeature), ctx, id, request)
}
