// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/kornharem08/app/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx
func (_m *IRepository) Find(ctx context.Context) ([]model.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmployeeId provides a mock function with given fields: ctx, employeeId
func (_m *IRepository) FindByEmployeeId(ctx context.Context, employeeId string) (*model.User, error) {
	ret := _m.Called(ctx, employeeId)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmployeeId")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.User, error)); ok {
		return rf(ctx, employeeId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, employeeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, employeeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
