// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "rental_easy.in/m/pkg/models"
)

// MockDataBase is a mock of DataBase interface.
type MockDataBase struct {
	ctrl     *gomock.Controller
	recorder *MockDataBaseMockRecorder
}

// MockDataBaseMockRecorder is the mock recorder for MockDataBase.
type MockDataBaseMockRecorder struct {
	mock *MockDataBase
}

// NewMockDataBase creates a new mock instance.
func NewMockDataBase(ctrl *gomock.Controller) *MockDataBase {
	mock := &MockDataBase{ctrl: ctrl}
	mock.recorder = &MockDataBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBase) EXPECT() *MockDataBaseMockRecorder {
	return m.recorder
}

// AddBooking mocks base method.
func (m *MockDataBase) AddBooking(arg0 models.Booking, arg1 models.Item) (models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBooking", arg0, arg1)
	ret0, _ := ret[0].(models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBooking indicates an expected call of AddBooking.
func (mr *MockDataBaseMockRecorder) AddBooking(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBooking", reflect.TypeOf((*MockDataBase)(nil).AddBooking), arg0, arg1)
}

// AddItem mocks base method.
func (m *MockDataBase) AddItem(arg0 models.Item) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddItem", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddItem indicates an expected call of AddItem.
func (mr *MockDataBaseMockRecorder) AddItem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItem", reflect.TypeOf((*MockDataBase)(nil).AddItem), arg0)
}

// AddReview mocks base method.
func (m *MockDataBase) AddReview(arg0 models.Review) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReview", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddReview indicates an expected call of AddReview.
func (mr *MockDataBaseMockRecorder) AddReview(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReview", reflect.TypeOf((*MockDataBase)(nil).AddReview), arg0)
}

// CreateUser mocks base method.
func (m *MockDataBase) CreateUser(arg0 models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDataBaseMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDataBase)(nil).CreateUser), arg0)
}

// DeleteItem mocks base method.
func (m *MockDataBase) DeleteItem(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockDataBaseMockRecorder) DeleteItem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockDataBase)(nil).DeleteItem), arg0)
}

// DeleteReview mocks base method.
func (m *MockDataBase) DeleteReview(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReview indicates an expected call of DeleteReview.
func (mr *MockDataBaseMockRecorder) DeleteReview(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockDataBase)(nil).DeleteReview), arg0)
}

// GetBookings mocks base method.
func (m *MockDataBase) GetBookings(arg0 int) ([]models.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookings", arg0)
	ret0, _ := ret[0].([]models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookings indicates an expected call of GetBookings.
func (mr *MockDataBaseMockRecorder) GetBookings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookings", reflect.TypeOf((*MockDataBase)(nil).GetBookings), arg0)
}

// GetItemById mocks base method.
func (m *MockDataBase) GetItemById(arg0 int) (models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemById", arg0)
	ret0, _ := ret[0].(models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemById indicates an expected call of GetItemById.
func (mr *MockDataBaseMockRecorder) GetItemById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemById", reflect.TypeOf((*MockDataBase)(nil).GetItemById), arg0)
}

// GetItemName mocks base method.
func (m *MockDataBase) GetItemName(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemName indicates an expected call of GetItemName.
func (mr *MockDataBaseMockRecorder) GetItemName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemName", reflect.TypeOf((*MockDataBase)(nil).GetItemName), arg0)
}

// GetItems mocks base method.
func (m *MockDataBase) GetItems() ([]models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems")
	ret0, _ := ret[0].([]models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockDataBaseMockRecorder) GetItems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockDataBase)(nil).GetItems))
}

// GetItemsofOwner mocks base method.
func (m *MockDataBase) GetItemsofOwner(arg0 int) ([]models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemsofOwner", arg0)
	ret0, _ := ret[0].([]models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemsofOwner indicates an expected call of GetItemsofOwner.
func (mr *MockDataBaseMockRecorder) GetItemsofOwner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemsofOwner", reflect.TypeOf((*MockDataBase)(nil).GetItemsofOwner), arg0)
}

// GetReviews mocks base method.
func (m *MockDataBase) GetReviews(arg0 int) ([]models.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviews", arg0)
	ret0, _ := ret[0].([]models.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviews indicates an expected call of GetReviews.
func (mr *MockDataBaseMockRecorder) GetReviews(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviews", reflect.TypeOf((*MockDataBase)(nil).GetReviews), arg0)
}

// GetUser mocks base method.
func (m *MockDataBase) GetUser(arg0 int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockDataBaseMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDataBase)(nil).GetUser), arg0)
}

// GetUserEmail mocks base method.
func (m *MockDataBase) GetUserEmail(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserEmail", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserEmail indicates an expected call of GetUserEmail.
func (mr *MockDataBaseMockRecorder) GetUserEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserEmail", reflect.TypeOf((*MockDataBase)(nil).GetUserEmail), arg0)
}

// SearchItems mocks base method.
func (m *MockDataBase) SearchItems(arg0, arg1 string) ([]models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchItems", arg0, arg1)
	ret0, _ := ret[0].([]models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchItems indicates an expected call of SearchItems.
func (mr *MockDataBaseMockRecorder) SearchItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchItems", reflect.TypeOf((*MockDataBase)(nil).SearchItems), arg0, arg1)
}

// UpdateItem mocks base method.
func (m *MockDataBase) UpdateItem(item models.Item) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", item)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockDataBaseMockRecorder) UpdateItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockDataBase)(nil).UpdateItem), item)
}

// UpdateReview mocks base method.
func (m *MockDataBase) UpdateReview(arg0 models.Review) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReview indicates an expected call of UpdateReview.
func (mr *MockDataBaseMockRecorder) UpdateReview(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockDataBase)(nil).UpdateReview), arg0)
}

// UpdateUser mocks base method.
func (m *MockDataBase) UpdateUser(arg0 models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockDataBaseMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockDataBase)(nil).UpdateUser), arg0)
}

// UserAlreadyAddedReview mocks base method.
func (m *MockDataBase) UserAlreadyAddedReview(arg0, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserAlreadyAddedReview", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserAlreadyAddedReview indicates an expected call of UserAlreadyAddedReview.
func (mr *MockDataBaseMockRecorder) UserAlreadyAddedReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserAlreadyAddedReview", reflect.TypeOf((*MockDataBase)(nil).UserAlreadyAddedReview), arg0, arg1)
}

// UserAlreadyBooked mocks base method.
func (m *MockDataBase) UserAlreadyBooked(arg0, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserAlreadyBooked", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserAlreadyBooked indicates an expected call of UserAlreadyBooked.
func (mr *MockDataBaseMockRecorder) UserAlreadyBooked(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserAlreadyBooked", reflect.TypeOf((*MockDataBase)(nil).UserAlreadyBooked), arg0, arg1)
}
