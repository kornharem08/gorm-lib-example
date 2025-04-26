package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kornharem08/app/internal/users"
	sqlwrap "github.com/kornharem08/gorm"
)

type IHandler struct {
	UserService users.IService
}

func NewHandler(dbconn sqlwrap.ISQLConnect) *IHandler {
	return &IHandler{
		UserService: users.NewService(dbconn),
	}
}

func (h IHandler) Find(c *gin.Context) {
	users, err := h.UserService.Find(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h IHandler) FindByEmployeeId(c *gin.Context) {
	employeeId := c.Param("employeeId")
	if employeeId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee ID is required"})
		return
	}

	user, err := h.UserService.FindByEmployeeId(c.Request.Context(), employeeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
