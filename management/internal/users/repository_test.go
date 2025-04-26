package users_test

import (
	"context"
	"testing"

	"github.com/kornharem08/app/internal/model"
	"github.com/kornharem08/app/internal/users"
	"github.com/kornharem08/gorm/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRepository_Find(t *testing.T) {
	// Create mocks
	mockSQL := new(mocks.ISQL)
	mockDB := new(mocks.IDatabase)
	mockConn := new(mocks.ISQLConnect)

	// Setup expectations
	mockConn.On("Database").Return(mockDB)
	mockDB.On("Table", mock.AnythingOfType("*model.User")).Return(mockSQL)

	// Success case
	t.Run("Success", func(t *testing.T) {
		// Setup mock to return success and populate users
		mockSQL.On("Find", mock.Anything, mock.AnythingOfType("*[]model.User")).
			Run(func(args mock.Arguments) {
				// Set test data in the users slice
				users := args.Get(1).(*[]model.User)
				*users = []model.User{
					{Contactid: "C0123456789", ADUserId: "AD123456789", EmployeeId: "E123456789", FirstName: "John", LastName: "Doe", LocalFirstName: "De", LocalLastName: "Dad", EmployeeTag: "CONTRACTOR"},
					{Contactid: "C0234567890", ADUserId: "AD234567890", EmployeeId: "E234567890", FirstName: "May", LastName: "Fah", LocalFirstName: "Ma", LocalLastName: "Loi", EmployeeTag: "STAFF"},
				}
			}).
			Return(nil).Once()

		// Create repository with mock
		repo := users.NewRepository(mockConn)

		// Call the method
		ctx := context.Background()
		result, err := repo.Find(ctx)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "C0123456789", result[0].Contactid)
		assert.Equal(t, "John", result[0].FirstName)

		mockSQL.AssertExpectations(t)
		mockDB.AssertExpectations(t)
		mockConn.AssertExpectations(t)
	})

	// Error case
	t.Run("Database Error", func(t *testing.T) {
		// Setup mock to return error
		mockSQL.On("Find", mock.Anything, mock.AnythingOfType("*[]model.User")).
			Return(assert.AnError).Once()

		// Create repository with mock
		repo := users.NewRepository(mockConn)

		// Call the method
		ctx := context.Background()
		result, err := repo.Find(ctx)

		// Assertions
		assert.Error(t, err)
		assert.Empty(t, result)

		mockSQL.AssertExpectations(t)
		mockDB.AssertExpectations(t)
		mockConn.AssertExpectations(t)
	})
}
