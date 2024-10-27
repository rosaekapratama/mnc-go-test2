// Code generated by mockery v2.46.3. DO NOT EDIT.

package v1

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// MockUserRestController is an autogenerated mock type for the UserRestController type
type MockUserRestController struct {
	mock.Mock
}

type MockUserRestController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRestController) EXPECT() *MockUserRestController_Expecter {
	return &MockUserRestController_Expecter{mock: &_m.Mock}
}

// Login provides a mock function with given fields: c
func (_m *MockUserRestController) Login(c *gin.Context) {
	_m.Called(c)
}

// MockUserRestController_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type MockUserRestController_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MockUserRestController_Expecter) Login(c interface{}) *MockUserRestController_Login_Call {
	return &MockUserRestController_Login_Call{Call: _e.mock.On("Login", c)}
}

func (_c *MockUserRestController_Login_Call) Run(run func(c *gin.Context)) *MockUserRestController_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockUserRestController_Login_Call) Return() *MockUserRestController_Login_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUserRestController_Login_Call) RunAndReturn(run func(*gin.Context)) *MockUserRestController_Login_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: c
func (_m *MockUserRestController) Register(c *gin.Context) {
	_m.Called(c)
}

// MockUserRestController_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockUserRestController_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - c *gin.Context
func (_e *MockUserRestController_Expecter) Register(c interface{}) *MockUserRestController_Register_Call {
	return &MockUserRestController_Register_Call{Call: _e.mock.On("Register", c)}
}

func (_c *MockUserRestController_Register_Call) Run(run func(c *gin.Context)) *MockUserRestController_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockUserRestController_Register_Call) Return() *MockUserRestController_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUserRestController_Register_Call) RunAndReturn(run func(*gin.Context)) *MockUserRestController_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRestController creates a new instance of MockUserRestController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRestController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRestController {
	mock := &MockUserRestController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}