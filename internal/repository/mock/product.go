// Code generated by MockGen. DO NOT EDIT.
// Source: product.go
//
// Generated by this command:
//
//	mockgen -source=product.go -destination=mock/product.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	model "github.com/binus-thesis-team/product-service/internal/model"
	pagination "github.com/binus-thesis-team/product-service/pkg/pagination"
	gomock "go.uber.org/mock/gomock"
)

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProduct) Create(ctx context.Context, product model.Product) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, product)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductMockRecorder) Create(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProduct)(nil).Create), ctx, product)
}

// Delete mocks base method.
func (m *MockProduct) Delete(ctx context.Context, productID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductMockRecorder) Delete(ctx, productID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProduct)(nil).Delete), ctx, productID)
}

// GetDetail mocks base method.
func (m *MockProduct) GetDetail(ctx context.Context, id int) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetail", ctx, id)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetail indicates an expected call of GetDetail.
func (mr *MockProductMockRecorder) GetDetail(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetail", reflect.TypeOf((*MockProduct)(nil).GetDetail), ctx, id)
}

// GetList mocks base method.
func (m *MockProduct) GetList(ctx context.Context, pagination pagination.Pagination) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, pagination)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockProductMockRecorder) GetList(ctx, pagination any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockProduct)(nil).GetList), ctx, pagination)
}

// Update mocks base method.
func (m *MockProduct) Update(ctx context.Context, product model.Product) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, product)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductMockRecorder) Update(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProduct)(nil).Update), ctx, product)
}
