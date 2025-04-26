package users_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/kornharem08/app/internal/model"
	"github.com/kornharem08/app/internal/users"
	"github.com/kornharem08/app/internal/users/mocks"

	sqlwrapMock "github.com/kornharem08/gorm/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewService(t *testing.T) {
	// Create mocks
	mockSQL := new(sqlwrapMock.ISQL)
	mockDB := new(sqlwrapMock.IDatabase)
	mockConn := new(sqlwrapMock.ISQLConnect)

	// Setup repo with mocks
	mockConn.On("Database").Return(mockDB)
	mockDB.On("Table", mock.AnythingOfType("*model.User")).Return(mockSQL)

	// Create repository with mock connection
	service := users.NewService(mockConn)
	v := reflect.Indirect(reflect.ValueOf(service))

	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}

}

func TestService_Find(t *testing.T) {
	var (
		mockRepository   *mocks.IRepository
		service          *users.Service
		ctx              context.Context
		expectedResponse []model.User
		mockError        error
	)

	beforeEach := func() {
		mockRepository = new(mocks.IRepository)
		mockError = nil
		ctx = context.Background()
		service = &users.Service{
			Repository: mockRepository,
		}
		expectedResponse = nil

		mockError = nil

		mockRepository.On("Find", context.Background()).Return(
			func(ctx context.Context) []model.User {
				return expectedResponse
			},
			func(ctx context.Context) error {
				return mockError
			},
		)

	}

	t.Run("Success", func(t *testing.T) {
		beforeEach()
		// Setup mock to return users
		expectedResponse = []model.User{
			{Contactid: "C0123456789", ADUserId: "AD123456789", EmployeeId: "E123456789", FirstName: "John", LastName: "Doe", LocalFirstName: "De", LocalLastName: "Dad", EmployeeTag: "CONTRACTOR"},
			{Contactid: "C0234567890", ADUserId: "AD234567890", EmployeeId: "E234567890", FirstName: "May", LastName: "Fah", LocalFirstName: "Ma", LocalLastName: "Loi", EmployeeTag: "STAFF"},
		}

		result, err := service.Find(ctx)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "C0123456789", result[0].Contactid)
		assert.Equal(t, "John", result[0].FirstName)

		mockRepository.AssertExpectations(t)
	})

	// Error case
	t.Run("Repository Error", func(t *testing.T) {
		beforeEach()
		mockError = errors.New("mock error")

		result, err := service.Find(ctx)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Equal(t, mockError, err)

		mockRepository.AssertExpectations(t)
	})
}

func TestService_FindByEmployeeId(t *testing.T) {
	var (
		mockRepository   *mocks.IRepository
		service          *users.Service
		ctx              context.Context
		expectedResponse *model.User
		mockError        error
		employeeId       string
	)

	beforeEach := func() {
		mockRepository = new(mocks.IRepository)
		mockError = nil
		ctx = context.Background()
		employeeId = "E123456789"
		service = &users.Service{
			Repository: mockRepository,
		}
		expectedResponse = nil

		mockRepository.On("FindByEmployeeId", context.Background(), employeeId).Return(
			func(ctx context.Context, empId string) *model.User {
				return expectedResponse
			},
			func(ctx context.Context, empId string) error {
				return mockError
			},
		)
	}

	t.Run("Success", func(t *testing.T) {
		beforeEach()
		// Setup mock to return user
		expectedResponse = &model.User{
			Contactid:      "C0123456789",
			ADUserId:       "AD123456789",
			EmployeeId:     "E123456789",
			FirstName:      "John",
			LastName:       "Doe",
			LocalFirstName: "De",
			LocalLastName:  "Dad",
			EmployeeTag:    "CONTRACTOR",
		}

		result, err := service.FindByEmployeeId(ctx, employeeId)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "C0123456789", result.Contactid)
		assert.Equal(t, "John", result.FirstName)
		assert.Equal(t, "E123456789", result.EmployeeId)

		mockRepository.AssertExpectations(t)
	})

	// Error case
	t.Run("Repository Error", func(t *testing.T) {
		beforeEach()
		mockError = errors.New("user not found")

		result, err := service.FindByEmployeeId(ctx, employeeId)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, mockError, err)

		mockRepository.AssertExpectations(t)
	})
}
