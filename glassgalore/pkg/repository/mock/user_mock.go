// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/user.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	domain "GlassGalore/pkg/domain"
	models "GlassGalore/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(id int, address models.AddAddress, result bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", id, address, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(id, address, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), id, address, result)
}

// ChangePassword mocks base method.
func (m *MockUserRepository) ChangePassword(id int, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", id, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepositoryMockRecorder) ChangePassword(id, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepository)(nil).ChangePassword), id, password)
}

// CheckIfFirstAddress mocks base method.
func (m *MockUserRepository) CheckIfFirstAddress(id int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfFirstAddress", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckIfFirstAddress indicates an expected call of CheckIfFirstAddress.
func (mr *MockUserRepositoryMockRecorder) CheckIfFirstAddress(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfFirstAddress", reflect.TypeOf((*MockUserRepository)(nil).CheckIfFirstAddress), id)
}

// CheckUserAvailability mocks base method.
func (m *MockUserRepository) CheckUserAvailability(email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserAvailability", email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUserAvailability indicates an expected call of CheckUserAvailability.
func (mr *MockUserRepositoryMockRecorder) CheckUserAvailability(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserAvailability", reflect.TypeOf((*MockUserRepository)(nil).CheckUserAvailability), email)
}

// EditDetails mocks base method.
func (m *MockUserRepository) EditDetails(id int, user models.EditDetailsResponse) (models.EditDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditDetails", id, user)
	ret0, _ := ret[0].(models.EditDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditDetails indicates an expected call of EditDetails.
func (mr *MockUserRepositoryMockRecorder) EditDetails(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDetails", reflect.TypeOf((*MockUserRepository)(nil).EditDetails), id, user)
}

// FindCartQuantity mocks base method.
func (m *MockUserRepository) FindCartQuantity(cart_id, inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCartQuantity", cart_id, inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCartQuantity indicates an expected call of FindCartQuantity.
func (mr *MockUserRepositoryMockRecorder) FindCartQuantity(cart_id, inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartQuantity", reflect.TypeOf((*MockUserRepository)(nil).FindCartQuantity), cart_id, inventory_id)
}

// FindCategory mocks base method.
func (m *MockUserRepository) FindCategory(inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategory", inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategory indicates an expected call of FindCategory.
func (mr *MockUserRepositoryMockRecorder) FindCategory(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategory", reflect.TypeOf((*MockUserRepository)(nil).FindCategory), inventory_id)
}

// FindPrice mocks base method.
func (m *MockUserRepository) FindPrice(inventory_id int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPrice", inventory_id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrice indicates an expected call of FindPrice.
func (mr *MockUserRepositoryMockRecorder) FindPrice(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPrice", reflect.TypeOf((*MockUserRepository)(nil).FindPrice), inventory_id)
}

// FindProductNames mocks base method.
func (m *MockUserRepository) FindProductNames(inventory_id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductNames", inventory_id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductNames indicates an expected call of FindProductNames.
func (mr *MockUserRepositoryMockRecorder) FindProductNames(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductNames", reflect.TypeOf((*MockUserRepository)(nil).FindProductNames), inventory_id)
}

// FindStock mocks base method.
func (m *MockUserRepository) FindStock(id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStock", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStock indicates an expected call of FindStock.
func (mr *MockUserRepositoryMockRecorder) FindStock(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStock", reflect.TypeOf((*MockUserRepository)(nil).FindStock), id)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", user)
	ret0, _ := ret[0].(models.UserSignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), user)
}

// GetAddresses mocks base method.
func (m *MockUserRepository) GetAddresses(id int) ([]domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses", id)
	ret0, _ := ret[0].([]domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockUserRepositoryMockRecorder) GetAddresses(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockUserRepository)(nil).GetAddresses), id)
}

// GetCartID mocks base method.
func (m *MockUserRepository) GetCartID(id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartID", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartID indicates an expected call of GetCartID.
func (mr *MockUserRepositoryMockRecorder) GetCartID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartID", reflect.TypeOf((*MockUserRepository)(nil).GetCartID), id)
}

// GetCatOfferr mocks base method.
func (m *MockUserRepository) GetCatOfferr(id int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCatOfferr", id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCatOfferr indicates an expected call of GetCatOfferr.
func (mr *MockUserRepositoryMockRecorder) GetCatOfferr(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCatOfferr", reflect.TypeOf((*MockUserRepository)(nil).GetCatOfferr), id)
}

// GetPassword mocks base method.
func (m *MockUserRepository) GetPassword(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockUserRepositoryMockRecorder) GetPassword(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockUserRepository)(nil).GetPassword), id)
}

// GetProductsInCart mocks base method.
func (m *MockUserRepository) GetProductsInCart(cart_id int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsInCart", cart_id)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsInCart indicates an expected call of GetProductsInCart.
func (mr *MockUserRepositoryMockRecorder) GetProductsInCart(cart_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsInCart", reflect.TypeOf((*MockUserRepository)(nil).GetProductsInCart), cart_id)
}

// GetUserDetails mocks base method.
func (m *MockUserRepository) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", id)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockUserRepositoryMockRecorder) GetUserDetails(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockUserRepository)(nil).GetUserDetails), id)
}

// RemoveFromCart mocks base method.
func (m *MockUserRepository) RemoveFromCart(cart, inventory int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromCart", cart, inventory)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromCart indicates an expected call of RemoveFromCart.
func (mr *MockUserRepositoryMockRecorder) RemoveFromCart(cart, inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromCart", reflect.TypeOf((*MockUserRepository)(nil).RemoveFromCart), cart, inventory)
}

// UpdateQuantity mocks base method.
func (m *MockUserRepository) UpdateQuantity(id, inv_id, qty int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantity", id, inv_id, qty)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuantity indicates an expected call of UpdateQuantity.
func (mr *MockUserRepositoryMockRecorder) UpdateQuantity(id, inv_id, qty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantity", reflect.TypeOf((*MockUserRepository)(nil).UpdateQuantity), id, inv_id, qty)
}

// UserBlockStatus mocks base method.
func (m *MockUserRepository) UserBlockStatus(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserBlockStatus", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserBlockStatus indicates an expected call of UserBlockStatus.
func (mr *MockUserRepositoryMockRecorder) UserBlockStatus(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserBlockStatus", reflect.TypeOf((*MockUserRepository)(nil).UserBlockStatus), email)
}

// UserSignUp mocks base method.
func (m *MockUserRepository) UserSignUp(arg0 models.UserDetails) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", arg0)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserRepositoryMockRecorder) UserSignUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserRepository)(nil).UserSignUp), arg0)
}
