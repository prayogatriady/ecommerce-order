// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/prayogatriady/ecommerce-module/model"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderRepository) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, order)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryMockRecorder) CreateOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrder), ctx, order)
}

// CreateOrderDetail mocks base method.
func (m *MockOrderRepository) CreateOrderDetail(ctx context.Context, orderDetail *[]model.OrderDetail) (*[]model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderDetail", ctx, orderDetail)
	ret0, _ := ret[0].(*[]model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderDetail indicates an expected call of CreateOrderDetail.
func (mr *MockOrderRepositoryMockRecorder) CreateOrderDetail(ctx, orderDetail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderDetail", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrderDetail), ctx, orderDetail)
}

// FindOrder mocks base method.
func (m *MockOrderRepository) FindOrder(ctx context.Context, orderId int64) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrder", ctx, orderId)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrder indicates an expected call of FindOrder.
func (mr *MockOrderRepositoryMockRecorder) FindOrder(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrder", reflect.TypeOf((*MockOrderRepository)(nil).FindOrder), ctx, orderId)
}

// FindOrderDetail mocks base method.
func (m *MockOrderRepository) FindOrderDetail(ctx context.Context, orderDetailId, orderId int64) (*model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderDetail", ctx, orderDetailId, orderId)
	ret0, _ := ret[0].(*model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrderDetail indicates an expected call of FindOrderDetail.
func (mr *MockOrderRepositoryMockRecorder) FindOrderDetail(ctx, orderDetailId, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderDetail", reflect.TypeOf((*MockOrderRepository)(nil).FindOrderDetail), ctx, orderDetailId, orderId)
}

// FindOrderDetails mocks base method.
func (m *MockOrderRepository) FindOrderDetails(ctx context.Context, orderId int64) (*[]model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderDetails", ctx, orderId)
	ret0, _ := ret[0].(*[]model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrderDetails indicates an expected call of FindOrderDetails.
func (mr *MockOrderRepositoryMockRecorder) FindOrderDetails(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderDetails", reflect.TypeOf((*MockOrderRepository)(nil).FindOrderDetails), ctx, orderId)
}

// FindOrders mocks base method.
func (m *MockOrderRepository) FindOrders(ctx context.Context) (*[]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrders", ctx)
	ret0, _ := ret[0].(*[]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrders indicates an expected call of FindOrders.
func (mr *MockOrderRepositoryMockRecorder) FindOrders(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrders", reflect.TypeOf((*MockOrderRepository)(nil).FindOrders), ctx)
}
