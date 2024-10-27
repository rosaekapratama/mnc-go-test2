package services

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/rosaekapratama/go-starter/mocks/database"
	"github.com/rosaekapratama/go-starter/response"
	"github.com/rosaekapratama/mnc-go-test2/crypto"
	"github.com/rosaekapratama/mnc-go-test2/mocks/repositories"
	repoModel "github.com/rosaekapratama/mnc-go-test2/models/repo"
	restModel "github.com/rosaekapratama/mnc-go-test2/models/rest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Register(t *testing.T) {
	tests := []struct {
		name       string
		input      *restModel.RegisterRequest
		setupMocks func(
			sqlDB sqlmock.Sqlmock,
			mockUserRepo *repositories.MockUserRepository,
			mockAccountRepo *repositories.MockAccountRepository,
		)
		expectedStatus string
		expectedMsg    string
		expectError    bool
	}{
		{
			name: "successfully register user",
			input: &restModel.RegisterRequest{
				PhoneNumber: "085722955088",
				Pin:         "123456",
				FirstName:   "John",
				LastName:    "Doe",
				Address:     "New York",
			},
			setupMocks: func(
				sqlDB sqlmock.Sqlmock,
				mockUserRepo *repositories.MockUserRepository,
				mockAccountRepo *repositories.MockAccountRepository,
			) {
				sqlDB.ExpectBegin()
				mockUserRepo.EXPECT().FindByPhoneNo(mockAny, "085722955088").Return(nil, nil)
				mockUserRepo.EXPECT().Save(mockAny, mockAny, mockAny).Return(nil)
				mockAccountRepo.EXPECT().CreateSaving(mockAny, mockAny, mockAny).Return(nil)
				sqlDB.ExpectCommit()
			},
			expectedStatus: "SUCCESS",
			expectedMsg:    "",
			expectError:    false,
		},
		{
			name: "phone number already exists",
			input: &restModel.RegisterRequest{
				PhoneNumber: "085722955088",
				Pin:         "123456",
			},
			setupMocks: func(
				sqlDB sqlmock.Sqlmock,
				mockUserRepo *repositories.MockUserRepository,
				mockAccountRepo *repositories.MockAccountRepository,
			) {
				mockUserRepo.EXPECT().FindByPhoneNo(mockAny, "085722955088").Return(&repoModel.User{}, nil)
			},
			expectedStatus: "",
			expectedMsg:    "Phone number already exists",
			expectError:    false,
		},
		{
			name: "phone number is empty",
			input: &restModel.RegisterRequest{
				PhoneNumber: "",
				Pin:         "123456",
			},
			setupMocks: func(
				sqlDB sqlmock.Sqlmock,
				mockUserRepo *repositories.MockUserRepository,
				mockAccountRepo *repositories.MockAccountRepository,
			) {
			},
			expectedStatus: "",
			expectedMsg:    "Phone number is empty",
			expectError:    false,
		},
		{
			name: "pin is empty",
			input: &restModel.RegisterRequest{
				PhoneNumber: "085722955088",
				Pin:         "",
			},
			setupMocks: func(
				sqlDB sqlmock.Sqlmock,
				mockUserRepo *repositories.MockUserRepository,
				mockAccountRepo *repositories.MockAccountRepository,
			) {
			},
			expectedStatus: "",
			expectedMsg:    "Pin is empty",
			expectError:    false,
		},
		{
			name: "error when saving user",
			input: &restModel.RegisterRequest{
				PhoneNumber: "085722955088",
				Pin:         "123456",
			},
			setupMocks: func(
				sqlDB sqlmock.Sqlmock,
				mockUserRepo *repositories.MockUserRepository,
				mockAccountRepo *repositories.MockAccountRepository,
			) {
				mockUserRepo.EXPECT().FindByPhoneNo(mockAny, "085722955088").Return(nil, nil)
				mockUserRepo.EXPECT().Save(mockAny, mockAny, mockAny).Return(errors.New("db error"))
			},
			expectedStatus: "",
			expectedMsg:    response.GeneralError.Description(),
			expectError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			ctx := context.Background()
			secret := "secret"

			// Create mock repositories and database transaction
			_, sqlDB, _ := database.SetupMockDB()
			mockUserRepo := new(repositories.MockUserRepository)
			mockAccountRepo := new(repositories.MockAccountRepository)
			userService := NewUserService(ctx, secret, mockUserRepo, mockAccountRepo)

			// Set up mocks based on test case
			test.setupMocks(sqlDB, mockUserRepo, mockAccountRepo)

			// Execute the service
			res, err := userService.Register(context.TODO(), test.input)

			// Assert based on the expected outcome
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedStatus, res.Status)
				assert.Equal(t, test.expectedMsg, res.Message)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	tests := []struct {
		name       string
		req        *restModel.LoginRequest
		setupMocks func(sqlDB sqlmock.Sqlmock, mockUserRepo *repositories.MockUserRepository)
		assertFunc func(res *restModel.BaseResponse[*restModel.LoginResponse], err error)
	}{
		{
			name: "successful login",
			req: &restModel.LoginRequest{
				PhoneNumber: "1234567890",
				Pin:         "1234",
			},
			setupMocks: func(sqlDB sqlmock.Sqlmock, mockUserRepo *repositories.MockUserRepository) {
				sqlDB.ExpectBegin()
				mockUserRepo.EXPECT().FindByPhoneNo(mockAny, "1234567890").Return(&repoModel.User{
					ID:          uuid.New(),
					PhoneNumber: "1234567890",
					Pin:         crypto.Hash("1234"), // assuming your Hash method is correct
				}, nil)
				mockUserRepo.EXPECT().Save(mockAny, mockAny, mockAny).Return(nil)
				sqlDB.ExpectCommit()
			},
			assertFunc: func(res *restModel.BaseResponse[*restModel.LoginResponse], err error) {
				assert.NoError(t, err)
				assert.Equal(t, "SUCCESS", res.Status)
				assert.Equal(t, "", res.Message)
				assert.Equal(t, true, res.Result.AccessToken != "")
			},
		},
		{
			name: "phone number is empty",
			req: &restModel.LoginRequest{
				PhoneNumber: "",
				Pin:         "1234",
			},
			setupMocks: func(sqlDB sqlmock.Sqlmock, mockUserRepo *repositories.MockUserRepository) {},
			assertFunc: func(res *restModel.BaseResponse[*restModel.LoginResponse], err error) {
				assert.NoError(t, err)
				assert.Equal(t, "", res.Status)
				assert.Equal(t, "Phone number is empty", res.Message)
			},
		},
		{
			name: "pin is empty",
			req: &restModel.LoginRequest{
				PhoneNumber: "1234567890",
				Pin:         "",
			},
			setupMocks: func(sqlDB sqlmock.Sqlmock, mockUserRepo *repositories.MockUserRepository) {},
			assertFunc: func(res *restModel.BaseResponse[*restModel.LoginResponse], err error) {
				assert.NoError(t, err)
				assert.Equal(t, "", res.Status)
				assert.Equal(t, "Pin is empty", res.Message)
			},
		},
		{
			name: "user not found",
			req: &restModel.LoginRequest{
				PhoneNumber: "1234567890",
				Pin:         "1234",
			},
			setupMocks: func(sqlDB sqlmock.Sqlmock, mockUserRepo *repositories.MockUserRepository) {
				mockUserRepo.EXPECT().FindByPhoneNo(mockAny, "1234567890").Return(nil, nil)
			},
			assertFunc: func(res *restModel.BaseResponse[*restModel.LoginResponse], err error) {
				assert.NoError(t, err)
				assert.Equal(t, "", res.Status)
				assert.Equal(t, "User not found", res.Message)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			secret := "secret"

			// Create mock repositories and database transaction
			_, sqlDB, _ := database.SetupMockDB()
			mockUserRepo := new(repositories.MockUserRepository)
			mockAccountRepo := new(repositories.MockAccountRepository)
			userService := NewUserService(ctx, secret, mockUserRepo, mockAccountRepo)

			// Set up mocks based on test case
			test.setupMocks(sqlDB, mockUserRepo)

			// Execute the service
			res, err := userService.Login(context.TODO(), test.req)

			// Assert based on the expected outcome
			test.assertFunc(res, err)
		})
	}
}
