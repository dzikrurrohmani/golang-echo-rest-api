// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
)

// MockRestoUsecase is a mock of Usecase interface.
type MockRestoUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockRestoUsecaseMockRecorder
}

// MockRestoUsecaseMockRecorder is the mock recorder for MockRestoUsecase.
type MockRestoUsecaseMockRecorder struct {
	mock *MockRestoUsecase
}

// NewMockRestoUsecase creates a new mock instance.
func NewMockRestoUsecase(ctrl *gomock.Controller) *MockRestoUsecase {
	mock := &MockRestoUsecase{ctrl: ctrl}
	mock.recorder = &MockRestoUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestoUsecase) EXPECT() *MockRestoUsecaseMockRecorder {
	return m.recorder
}

// CheckSession mocks base method.
func (m *MockRestoUsecase) CheckSession(ctx context.Context, sessionData model.UserSession) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSession", ctx, sessionData)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession.
func (mr *MockRestoUsecaseMockRecorder) CheckSession(ctx, sessionData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockRestoUsecase)(nil).CheckSession), ctx, sessionData)
}

// GetMenuList mocks base method.
func (m *MockRestoUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuList", ctx, menuType)
	ret0, _ := ret[0].([]model.MenuItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuList indicates an expected call of GetMenuList.
func (mr *MockRestoUsecaseMockRecorder) GetMenuList(ctx, menuType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuList", reflect.TypeOf((*MockRestoUsecase)(nil).GetMenuList), ctx, menuType)
}

// GetOrderInfo mocks base method.
func (m *MockRestoUsecase) GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderInfo", ctx, request)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderInfo indicates an expected call of GetOrderInfo.
func (mr *MockRestoUsecaseMockRecorder) GetOrderInfo(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderInfo", reflect.TypeOf((*MockRestoUsecase)(nil).GetOrderInfo), ctx, request)
}

// Login mocks base method.
func (m *MockRestoUsecase) Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, request)
	ret0, _ := ret[0].(model.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockRestoUsecaseMockRecorder) Login(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRestoUsecase)(nil).Login), ctx, request)
}

// Order mocks base method.
func (m *MockRestoUsecase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order", ctx, request)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Order indicates an expected call of Order.
func (mr *MockRestoUsecaseMockRecorder) Order(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockRestoUsecase)(nil).Order), ctx, request)
}

// RegisterUser mocks base method.
func (m *MockRestoUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, request)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockRestoUsecaseMockRecorder) RegisterUser(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockRestoUsecase)(nil).RegisterUser), ctx, request)
}
