// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	entity "github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockUser) DeleteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockUserMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockUser)(nil).DeleteByID), ctx, ID)
}

// GetByID mocks base method.
func (m *MockUser) GetByID(ctx context.Context, ID string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserMockRecorder) GetByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUser)(nil).GetByID), ctx, ID)
}

// List mocks base method.
func (m *MockUser) List(ctx context.Context, email string) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, email)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserMockRecorder) List(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUser)(nil).List), ctx, email)
}

// Login mocks base method.
func (m *MockUser) Login(ctx context.Context, user *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserMockRecorder) Login(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUser)(nil).Login), ctx, user)
}

// LostPassword mocks base method.
func (m *MockUser) LostPassword(ctx context.Context, email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LostPassword", ctx, email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LostPassword indicates an expected call of LostPassword.
func (mr *MockUserMockRecorder) LostPassword(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LostPassword", reflect.TypeOf((*MockUser)(nil).LostPassword), ctx, email)
}

// ModifyByID mocks base method.
func (m *MockUser) ModifyByID(ctx context.Context, ID string, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyByID", ctx, ID, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyByID indicates an expected call of ModifyByID.
func (mr *MockUserMockRecorder) ModifyByID(ctx, ID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyByID", reflect.TypeOf((*MockUser)(nil).ModifyByID), ctx, ID, user)
}

// ResetPassword mocks base method.
func (m *MockUser) ResetPassword(ctx context.Context, ID, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", ctx, ID, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserMockRecorder) ResetPassword(ctx, ID, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUser)(nil).ResetPassword), ctx, ID, newPassword)
}

// Store mocks base method.
func (m *MockUser) Store(ctx context.Context, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockUserMockRecorder) Store(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUser)(nil).Store), ctx, user)
}

// MockPasswordAuth is a mock of PasswordAuth interface.
type MockPasswordAuth struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordAuthMockRecorder
}

// MockPasswordAuthMockRecorder is the mock recorder for MockPasswordAuth.
type MockPasswordAuthMockRecorder struct {
	mock *MockPasswordAuth
}

// NewMockPasswordAuth creates a new mock instance.
func NewMockPasswordAuth(ctrl *gomock.Controller) *MockPasswordAuth {
	mock := &MockPasswordAuth{ctrl: ctrl}
	mock.recorder = &MockPasswordAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordAuth) EXPECT() *MockPasswordAuthMockRecorder {
	return m.recorder
}

// HashAndSalt mocks base method.
func (m *MockPasswordAuth) HashAndSalt(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashAndSalt", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashAndSalt indicates an expected call of HashAndSalt.
func (mr *MockPasswordAuthMockRecorder) HashAndSalt(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashAndSalt", reflect.TypeOf((*MockPasswordAuth)(nil).HashAndSalt), password)
}

// Verify mocks base method.
func (m *MockPasswordAuth) Verify(hashed, password string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", hashed, password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockPasswordAuthMockRecorder) Verify(hashed, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockPasswordAuth)(nil).Verify), hashed, password)
}

// MockBarberShop is a mock of BarberShop interface.
type MockBarberShop struct {
	ctrl     *gomock.Controller
	recorder *MockBarberShopMockRecorder
}

// MockBarberShopMockRecorder is the mock recorder for MockBarberShop.
type MockBarberShopMockRecorder struct {
	mock *MockBarberShop
}

// NewMockBarberShop creates a new mock instance.
func NewMockBarberShop(ctrl *gomock.Controller) *MockBarberShop {
	mock := &MockBarberShop{ctrl: ctrl}
	mock.recorder = &MockBarberShopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBarberShop) EXPECT() *MockBarberShopMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockBarberShop) DeleteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockBarberShopMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockBarberShop)(nil).DeleteByID), ctx, ID)
}

// Find mocks base method.
func (m *MockBarberShop) Find(ctx context.Context, lat, lon, name, radius string) ([]*entity.BarberShop, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, lat, lon, name, radius)
	ret0, _ := ret[0].([]*entity.BarberShop)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockBarberShopMockRecorder) Find(ctx, lat, lon, name, radius interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockBarberShop)(nil).Find), ctx, lat, lon, name, radius)
}

// GetByID mocks base method.
func (m *MockBarberShop) GetByID(ctx context.Context, viewerID, ID string) (*entity.BarberShop, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, viewerID, ID)
	ret0, _ := ret[0].(*entity.BarberShop)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockBarberShopMockRecorder) GetByID(ctx, viewerID, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockBarberShop)(nil).GetByID), ctx, viewerID, ID)
}

// ModifyByID mocks base method.
func (m *MockBarberShop) ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyByID", ctx, ID, shop)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyByID indicates an expected call of ModifyByID.
func (mr *MockBarberShopMockRecorder) ModifyByID(ctx, ID, shop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyByID", reflect.TypeOf((*MockBarberShop)(nil).ModifyByID), ctx, ID, shop)
}

// Store mocks base method.
func (m *MockBarberShop) Store(ctx context.Context, shop *entity.BarberShop) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, shop)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockBarberShopMockRecorder) Store(ctx, shop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockBarberShop)(nil).Store), ctx, shop)
}

// MockCalendar is a mock of Calendar interface.
type MockCalendar struct {
	ctrl     *gomock.Controller
	recorder *MockCalendarMockRecorder
}

// MockCalendarMockRecorder is the mock recorder for MockCalendar.
type MockCalendarMockRecorder struct {
	mock *MockCalendar
}

// NewMockCalendar creates a new mock instance.
func NewMockCalendar(ctrl *gomock.Controller) *MockCalendar {
	mock := &MockCalendar{ctrl: ctrl}
	mock.recorder = &MockCalendarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalendar) EXPECT() *MockCalendarMockRecorder {
	return m.recorder
}

// GetByBarberShopID mocks base method.
func (m *MockCalendar) GetByBarberShopID(ctx context.Context, ID string) (*entity.Calendar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBarberShopID", ctx, ID)
	ret0, _ := ret[0].(*entity.Calendar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBarberShopID indicates an expected call of GetByBarberShopID.
func (mr *MockCalendarMockRecorder) GetByBarberShopID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBarberShopID", reflect.TypeOf((*MockCalendar)(nil).GetByBarberShopID), ctx, ID)
}

// MockAppointment is a mock of Appointment interface.
type MockAppointment struct {
	ctrl     *gomock.Controller
	recorder *MockAppointmentMockRecorder
}

// MockAppointmentMockRecorder is the mock recorder for MockAppointment.
type MockAppointmentMockRecorder struct {
	mock *MockAppointment
}

// NewMockAppointment creates a new mock instance.
func NewMockAppointment(ctrl *gomock.Controller) *MockAppointment {
	mock := &MockAppointment{ctrl: ctrl}
	mock.recorder = &MockAppointmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppointment) EXPECT() *MockAppointmentMockRecorder {
	return m.recorder
}

// Book mocks base method.
func (m *MockAppointment) Book(ctx context.Context, appointment *entity.Appointment) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Book", ctx, appointment)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Book indicates an expected call of Book.
func (mr *MockAppointmentMockRecorder) Book(ctx, appointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Book", reflect.TypeOf((*MockAppointment)(nil).Book), ctx, appointment)
}

// Cancel mocks base method.
func (m *MockAppointment) Cancel(ctx context.Context, appointment *entity.Appointment) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", ctx, appointment)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cancel indicates an expected call of Cancel.
func (mr *MockAppointmentMockRecorder) Cancel(ctx, appointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockAppointment)(nil).Cancel), ctx, appointment)
}

// DeleteByID mocks base method.
func (m *MockAppointment) DeleteByID(ctx context.Context, ID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockAppointmentMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockAppointment)(nil).DeleteByID), ctx, ID)
}

// MockHoliday is a mock of Holiday interface.
type MockHoliday struct {
	ctrl     *gomock.Controller
	recorder *MockHolidayMockRecorder
}

// MockHolidayMockRecorder is the mock recorder for MockHoliday.
type MockHolidayMockRecorder struct {
	mock *MockHoliday
}

// NewMockHoliday creates a new mock instance.
func NewMockHoliday(ctrl *gomock.Controller) *MockHoliday {
	mock := &MockHoliday{ctrl: ctrl}
	mock.recorder = &MockHolidayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHoliday) EXPECT() *MockHolidayMockRecorder {
	return m.recorder
}

// Set mocks base method.
func (m *MockHoliday) Set(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, shopID, date, unavailableEmployees)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockHolidayMockRecorder) Set(ctx, shopID, date, unavailableEmployees interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockHoliday)(nil).Set), ctx, shopID, date, unavailableEmployees)
}

// MockReview is a mock of Review interface.
type MockReview struct {
	ctrl     *gomock.Controller
	recorder *MockReviewMockRecorder
}

// MockReviewMockRecorder is the mock recorder for MockReview.
type MockReviewMockRecorder struct {
	mock *MockReview
}

// NewMockReview creates a new mock instance.
func NewMockReview(ctrl *gomock.Controller) *MockReview {
	mock := &MockReview{ctrl: ctrl}
	mock.recorder = &MockReviewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReview) EXPECT() *MockReviewMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockReview) DeleteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockReviewMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockReview)(nil).DeleteByID), ctx, ID)
}

// GetByBarberShop mocks base method.
func (m *MockReview) GetByBarberShop(ctx context.Context, shopID string) ([]*entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBarberShop", ctx, shopID)
	ret0, _ := ret[0].([]*entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBarberShop indicates an expected call of GetByBarberShop.
func (mr *MockReviewMockRecorder) GetByBarberShop(ctx, shopID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBarberShop", reflect.TypeOf((*MockReview)(nil).GetByBarberShop), ctx, shopID)
}

// Store mocks base method.
func (m *MockReview) Store(ctx context.Context, userID, shopID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, userID, shopID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockReviewMockRecorder) Store(ctx, userID, shopID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockReview)(nil).Store), ctx, userID, shopID)
}

// VoteByID mocks base method.
func (m *MockReview) VoteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VoteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// VoteByID indicates an expected call of VoteByID.
func (mr *MockReviewMockRecorder) VoteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VoteByID", reflect.TypeOf((*MockReview)(nil).VoteByID), ctx, ID)
}

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockUserRepo) DeleteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockUserRepoMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockUserRepo)(nil).DeleteByID), ctx, ID)
}

// GetByEmail mocks base method.
func (m *MockUserRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserRepoMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserRepo)(nil).GetByEmail), ctx, email)
}

// GetByID mocks base method.
func (m *MockUserRepo) GetByID(ctx context.Context, ID string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserRepoMockRecorder) GetByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepo)(nil).GetByID), ctx, ID)
}

// List mocks base method.
func (m *MockUserRepo) List(ctx context.Context, email string) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, email)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserRepoMockRecorder) List(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserRepo)(nil).List), ctx, email)
}

// ModifyByID mocks base method.
func (m *MockUserRepo) ModifyByID(ctx context.Context, ID string, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyByID", ctx, ID, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyByID indicates an expected call of ModifyByID.
func (mr *MockUserRepoMockRecorder) ModifyByID(ctx, ID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyByID", reflect.TypeOf((*MockUserRepo)(nil).ModifyByID), ctx, ID, user)
}

// Store mocks base method.
func (m *MockUserRepo) Store(ctx context.Context, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockUserRepoMockRecorder) Store(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUserRepo)(nil).Store), ctx, user)
}

// MockBarberShopRepo is a mock of BarberShopRepo interface.
type MockBarberShopRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBarberShopRepoMockRecorder
}

// MockBarberShopRepoMockRecorder is the mock recorder for MockBarberShopRepo.
type MockBarberShopRepoMockRecorder struct {
	mock *MockBarberShopRepo
}

// NewMockBarberShopRepo creates a new mock instance.
func NewMockBarberShopRepo(ctrl *gomock.Controller) *MockBarberShopRepo {
	mock := &MockBarberShopRepo{ctrl: ctrl}
	mock.recorder = &MockBarberShopRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBarberShopRepo) EXPECT() *MockBarberShopRepoMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockBarberShopRepo) DeleteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockBarberShopRepoMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockBarberShopRepo)(nil).DeleteByID), ctx, ID)
}

// Find mocks base method.
func (m *MockBarberShopRepo) Find(ctx context.Context, lat, lon, name, radius string) ([]*entity.BarberShop, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, lat, lon, name, radius)
	ret0, _ := ret[0].([]*entity.BarberShop)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockBarberShopRepoMockRecorder) Find(ctx, lat, lon, name, radius interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockBarberShopRepo)(nil).Find), ctx, lat, lon, name, radius)
}

// GetByID mocks base method.
func (m *MockBarberShopRepo) GetByID(ctx context.Context, ID string) (*entity.BarberShop, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ID)
	ret0, _ := ret[0].(*entity.BarberShop)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockBarberShopRepoMockRecorder) GetByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockBarberShopRepo)(nil).GetByID), ctx, ID)
}

// ModifyByID mocks base method.
func (m *MockBarberShopRepo) ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyByID", ctx, ID, shop)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyByID indicates an expected call of ModifyByID.
func (mr *MockBarberShopRepoMockRecorder) ModifyByID(ctx, ID, shop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyByID", reflect.TypeOf((*MockBarberShopRepo)(nil).ModifyByID), ctx, ID, shop)
}

// Store mocks base method.
func (m *MockBarberShopRepo) Store(ctx context.Context, shop *entity.BarberShop) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, shop)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockBarberShopRepoMockRecorder) Store(ctx, shop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockBarberShopRepo)(nil).Store), ctx, shop)
}

// MockSlotRepo is a mock of SlotRepo interface.
type MockSlotRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSlotRepoMockRecorder
}

// MockSlotRepoMockRecorder is the mock recorder for MockSlotRepo.
type MockSlotRepoMockRecorder struct {
	mock *MockSlotRepo
}

// NewMockSlotRepo creates a new mock instance.
func NewMockSlotRepo(ctrl *gomock.Controller) *MockSlotRepo {
	mock := &MockSlotRepo{ctrl: ctrl}
	mock.recorder = &MockSlotRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlotRepo) EXPECT() *MockSlotRepoMockRecorder {
	return m.recorder
}

// Book mocks base method.
func (m *MockSlotRepo) Book(ctx context.Context, appointment *entity.Appointment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Book", ctx, appointment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Book indicates an expected call of Book.
func (mr *MockSlotRepoMockRecorder) Book(ctx, appointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Book", reflect.TypeOf((*MockSlotRepo)(nil).Book), ctx, appointment)
}

// Cancel mocks base method.
func (m *MockSlotRepo) Cancel(ctx context.Context, appointment *entity.Appointment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", ctx, appointment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockSlotRepoMockRecorder) Cancel(ctx, appointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockSlotRepo)(nil).Cancel), ctx, appointment)
}

// GetByBarberShopID mocks base method.
func (m *MockSlotRepo) GetByBarberShopID(ctx context.Context, ID string) ([]*entity.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBarberShopID", ctx, ID)
	ret0, _ := ret[0].([]*entity.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBarberShopID indicates an expected call of GetByBarberShopID.
func (mr *MockSlotRepoMockRecorder) GetByBarberShopID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBarberShopID", reflect.TypeOf((*MockSlotRepo)(nil).GetByBarberShopID), ctx, ID)
}

// SetHoliday mocks base method.
func (m *MockSlotRepo) SetHoliday(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHoliday", ctx, shopID, date, unavailableEmployees)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHoliday indicates an expected call of SetHoliday.
func (mr *MockSlotRepoMockRecorder) SetHoliday(ctx, shopID, date, unavailableEmployees interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHoliday", reflect.TypeOf((*MockSlotRepo)(nil).SetHoliday), ctx, shopID, date, unavailableEmployees)
}

// MockShopViewRepo is a mock of ShopViewRepo interface.
type MockShopViewRepo struct {
	ctrl     *gomock.Controller
	recorder *MockShopViewRepoMockRecorder
}

// MockShopViewRepoMockRecorder is the mock recorder for MockShopViewRepo.
type MockShopViewRepoMockRecorder struct {
	mock *MockShopViewRepo
}

// NewMockShopViewRepo creates a new mock instance.
func NewMockShopViewRepo(ctrl *gomock.Controller) *MockShopViewRepo {
	mock := &MockShopViewRepo{ctrl: ctrl}
	mock.recorder = &MockShopViewRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShopViewRepo) EXPECT() *MockShopViewRepoMockRecorder {
	return m.recorder
}

// Store mocks base method.
func (m *MockShopViewRepo) Store(ctx context.Context, view *entity.ShopView) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, view)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockShopViewRepoMockRecorder) Store(ctx, view interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockShopViewRepo)(nil).Store), ctx, view)
}

// MockAppointmentRepo is a mock of AppointmentRepo interface.
type MockAppointmentRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAppointmentRepoMockRecorder
}

// MockAppointmentRepoMockRecorder is the mock recorder for MockAppointmentRepo.
type MockAppointmentRepoMockRecorder struct {
	mock *MockAppointmentRepo
}

// NewMockAppointmentRepo creates a new mock instance.
func NewMockAppointmentRepo(ctrl *gomock.Controller) *MockAppointmentRepo {
	mock := &MockAppointmentRepo{ctrl: ctrl}
	mock.recorder = &MockAppointmentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppointmentRepo) EXPECT() *MockAppointmentRepoMockRecorder {
	return m.recorder
}

// Book mocks base method.
func (m *MockAppointmentRepo) Book(ctx context.Context, appointment *entity.Appointment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Book", ctx, appointment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Book indicates an expected call of Book.
func (mr *MockAppointmentRepoMockRecorder) Book(ctx, appointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Book", reflect.TypeOf((*MockAppointmentRepo)(nil).Book), ctx, appointment)
}

// Cancel mocks base method.
func (m *MockAppointmentRepo) Cancel(ctx context.Context, appointment *entity.Appointment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", ctx, appointment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockAppointmentRepoMockRecorder) Cancel(ctx, appointment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockAppointmentRepo)(nil).Cancel), ctx, appointment)
}

// DeleteByID mocks base method.
func (m *MockAppointmentRepo) DeleteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockAppointmentRepoMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockAppointmentRepo)(nil).DeleteByID), ctx, ID)
}

// MockReviewRepo is a mock of ReviewRepo interface.
type MockReviewRepo struct {
	ctrl     *gomock.Controller
	recorder *MockReviewRepoMockRecorder
}

// MockReviewRepoMockRecorder is the mock recorder for MockReviewRepo.
type MockReviewRepoMockRecorder struct {
	mock *MockReviewRepo
}

// NewMockReviewRepo creates a new mock instance.
func NewMockReviewRepo(ctrl *gomock.Controller) *MockReviewRepo {
	mock := &MockReviewRepo{ctrl: ctrl}
	mock.recorder = &MockReviewRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewRepo) EXPECT() *MockReviewRepoMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockReviewRepo) DeleteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockReviewRepoMockRecorder) DeleteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockReviewRepo)(nil).DeleteByID), ctx, ID)
}

// GetByBarberShop mocks base method.
func (m *MockReviewRepo) GetByBarberShop(ctx context.Context, shopID string) ([]*entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBarberShop", ctx, shopID)
	ret0, _ := ret[0].([]*entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBarberShop indicates an expected call of GetByBarberShop.
func (mr *MockReviewRepoMockRecorder) GetByBarberShop(ctx, shopID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBarberShop", reflect.TypeOf((*MockReviewRepo)(nil).GetByBarberShop), ctx, shopID)
}

// Store mocks base method.
func (m *MockReviewRepo) Store(ctx context.Context, userID, shopID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, userID, shopID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockReviewRepoMockRecorder) Store(ctx, userID, shopID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockReviewRepo)(nil).Store), ctx, userID, shopID)
}

// VoteByID mocks base method.
func (m *MockReviewRepo) VoteByID(ctx context.Context, ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VoteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// VoteByID indicates an expected call of VoteByID.
func (mr *MockReviewRepoMockRecorder) VoteByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VoteByID", reflect.TypeOf((*MockReviewRepo)(nil).VoteByID), ctx, ID)
}

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}
