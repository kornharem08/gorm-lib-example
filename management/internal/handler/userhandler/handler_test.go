package userhandler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kornharem08/app/internal/handler/userhandler"
	"github.com/kornharem08/app/internal/model"
	"github.com/kornharem08/app/internal/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Find(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create mock service
	mockService := new(mocks.IService)

	// Create handler with mock service
	handler := &userhandler.IHandler{
		UserService: mockService,
	}

	// Success case
	t.Run("Success", func(t *testing.T) {
		// Setup mock to return users
		mockUsers := []model.User{
			{ID: 1, Name: "Test User", Email: "test@example.com"},
			{ID: 2, Name: "Another User", Email: "another@example.com"},
		}
		mockService.On("Find", mock.Anything).Return(mockUsers, nil).Once()

		// Create test context
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Call handler
		handler.Find(c)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse response
		var response map[string][]model.User
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Check data
		assert.Len(t, response["data"], 2)
		assert.Equal(t, uint(1), response["data"][0].ID)
		assert.Equal(t, "Test User", response["data"][0].Name)

		mockService.AssertExpectations(t)
	})

	// Error case
	t.Run("Service Error", func(t *testing.T) {
		// Setup mock to return error
		mockService.On("Find", mock.Anything).Return(nil, errors.New("service error")).Once()

		// Create test context
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Call handler
		handler.Find(c)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Parse response
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Check error message
		assert.Equal(t, "service error", response["error"])

		mockService.AssertExpectations(t)
	})
}
