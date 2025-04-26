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
