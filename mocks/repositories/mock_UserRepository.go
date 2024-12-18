// Code generated by mockery v2.46.3. DO NOT EDIT.

package repositories

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	repo "github.com/rosaekapratama/mnc-go-test2/models/repo"

	uuid "github.com/google/uuid"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// FindById provides a mock function with given fields: ctx, userId
func (_m *MockUserRepository) FindById(ctx context.Context, userId uuid.UUID) (*repo.User, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *repo.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*repo.User, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *repo.User); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repo.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type MockUserRepository_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - ctx context.Context
//   - userId uuid.UUID
func (_e *MockUserRepository_Expecter) FindById(ctx interface{}, userId interface{}) *MockUserRepository_FindById_Call {
	return &MockUserRepository_FindById_Call{Call: _e.mock.On("FindById", ctx, userId)}
}

func (_c *MockUserRepository_FindById_Call) Run(run func(ctx context.Context, userId uuid.UUID)) *MockUserRepository_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockUserRepository_FindById_Call) Return(user *repo.User, err error) *MockUserRepository_FindById_Call {
	_c.Call.Return(user, err)
	return _c
}

func (_c *MockUserRepository_FindById_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*repo.User, error)) *MockUserRepository_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// FindByPhoneNo provides a mock function with given fields: ctx, phoneNo
func (_m *MockUserRepository) FindByPhoneNo(ctx context.Context, phoneNo string) (*repo.User, error) {
	ret := _m.Called(ctx, phoneNo)

	if len(ret) == 0 {
		panic("no return value specified for FindByPhoneNo")
	}

	var r0 *repo.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*repo.User, error)); ok {
		return rf(ctx, phoneNo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *repo.User); ok {
		r0 = rf(ctx, phoneNo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repo.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, phoneNo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_FindByPhoneNo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByPhoneNo'
type MockUserRepository_FindByPhoneNo_Call struct {
	*mock.Call
}

// FindByPhoneNo is a helper method to define mock.On call
//   - ctx context.Context
//   - phoneNo string
func (_e *MockUserRepository_Expecter) FindByPhoneNo(ctx interface{}, phoneNo interface{}) *MockUserRepository_FindByPhoneNo_Call {
	return &MockUserRepository_FindByPhoneNo_Call{Call: _e.mock.On("FindByPhoneNo", ctx, phoneNo)}
}

func (_c *MockUserRepository_FindByPhoneNo_Call) Run(run func(ctx context.Context, phoneNo string)) *MockUserRepository_FindByPhoneNo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepository_FindByPhoneNo_Call) Return(user *repo.User, err error) *MockUserRepository_FindByPhoneNo_Call {
	_c.Call.Return(user, err)
	return _c
}

func (_c *MockUserRepository_FindByPhoneNo_Call) RunAndReturn(run func(context.Context, string) (*repo.User, error)) *MockUserRepository_FindByPhoneNo_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, tx, user
func (_m *MockUserRepository) Save(ctx context.Context, tx *gorm.DB, user *repo.User) error {
	ret := _m.Called(ctx, tx, user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *repo.User) error); ok {
		r0 = rf(ctx, tx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockUserRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *gorm.DB
//   - user *repo.User
func (_e *MockUserRepository_Expecter) Save(ctx interface{}, tx interface{}, user interface{}) *MockUserRepository_Save_Call {
	return &MockUserRepository_Save_Call{Call: _e.mock.On("Save", ctx, tx, user)}
}

func (_c *MockUserRepository_Save_Call) Run(run func(ctx context.Context, tx *gorm.DB, user *repo.User)) *MockUserRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*gorm.DB), args[2].(*repo.User))
	})
	return _c
}

func (_c *MockUserRepository_Save_Call) Return(err error) *MockUserRepository_Save_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUserRepository_Save_Call) RunAndReturn(run func(context.Context, *gorm.DB, *repo.User) error) *MockUserRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
