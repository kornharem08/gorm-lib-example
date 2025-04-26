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
			{ID: 1, Name: "Test User", Email: "test@example.com"},
			{ID: 2, Name: "Another User", Email: "another@example.com"},
		}

		result, err := service.Find(ctx)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, uint(1), result[0].ID)
		assert.Equal(t, "Test User", result[0].Name)

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
