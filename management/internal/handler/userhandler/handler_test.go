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

func setupTest() (*mocks.IService, *userhandler.IHandler, *gin.Context, *httptest.ResponseRecorder) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create mock service
	mockService := new(mocks.IService)

	// Create handler with mock service
	handler := &userhandler.IHandler{
		UserService: mockService,
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a test context with a request
	req, _ := http.NewRequest("GET", "/", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	return mockService, handler, c, w
}

func TestHandler_Find(t *testing.T) {
	// Success case
	t.Run("Success", func(t *testing.T) {
		mockService, handler, c, w := setupTest()

		// Setup mock to return users
		mockUsers := []model.User{
			{Contactid: "C0123456789", ADUserId: "AD123456789", EmployeeId: "E123456789", FirstName: "John", LastName: "Doe", LocalFirstName: "De", LocalLastName: "Dad", EmployeeTag: "CONTRACTOR"},
			{Contactid: "C0234567890", ADUserId: "AD234567890", EmployeeId: "E234567890", FirstName: "May", LastName: "Fah", LocalFirstName: "Ma", LocalLastName: "Loi", EmployeeTag: "STAFF"},
		}
		mockService.On("Find", mock.Anything).Return(mockUsers, nil).Once()

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
		assert.Equal(t, "C0123456789", response["data"][0].Contactid)
		assert.Equal(t, "John", response["data"][0].FirstName)

		mockService.AssertExpectations(t)
	})

	// Error case
	t.Run("Service Error", func(t *testing.T) {
		mockService, handler, c, w := setupTest()

		// Setup mock to return error
		mockService.On("Find", mock.Anything).Return(nil, errors.New("service error")).Once()

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

func TestHandler_FindByEmployeeId(t *testing.T) {
	// Success case
	t.Run("Success", func(t *testing.T) {
		mockService, handler, c, w := setupTest()

		// Setup mock to return user
		employeeId := "E123456789"
		mockUser := &model.User{
			Contactid:      "C0123456789",
			ADUserId:       "AD123456789",
			EmployeeId:     "E123456789",
			FirstName:      "John",
			LastName:       "Doe",
			LocalFirstName: "De",
			LocalLastName:  "Dad",
			EmployeeTag:    "CONTRACTOR",
		}
		mockService.On("FindByEmployeeId", mock.Anything, employeeId).Return(mockUser, nil).Once()

		// Setup request parameters
		c.Params = []gin.Param{
			{Key: "employeeId", Value: employeeId},
		}

		// Call handler
		handler.FindByEmployeeId(c)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse response
		var response map[string]*model.User
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Check data
		assert.NotNil(t, response["data"])
		assert.Equal(t, "C0123456789", response["data"].Contactid)
		assert.Equal(t, "John", response["data"].FirstName)
		assert.Equal(t, "E123456789", response["data"].EmployeeId)

		mockService.AssertExpectations(t)
	})

	// Error case - missing parameter
	t.Run("Missing EmployeeId", func(t *testing.T) {
		_, handler, c, w := setupTest()

		// Call handler with no params set
		handler.FindByEmployeeId(c)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Parse response
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Check error message
		assert.Equal(t, "employee ID is required", response["error"])
	})

	// Error case - service error
	t.Run("Service Error", func(t *testing.T) {
		mockService, handler, c, w := setupTest()

		// Setup mock to return error
		employeeId := "E123456789"
		mockService.On("FindByEmployeeId", mock.Anything, employeeId).Return(nil, errors.New("service error")).Once()

		// Setup request parameters
		c.Params = []gin.Param{
			{Key: "employeeId", Value: employeeId},
		}

		// Call handler
		handler.FindByEmployeeId(c)

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
