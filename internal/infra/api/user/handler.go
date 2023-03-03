package user

import (
	"net/http"

	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type (
	UserHandler struct {
		UserService ports.UserService
	}
)

func NewUserHandler(userService ports.UserService) *UserHandler {

	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	user := &entity.User{}

	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully :D"})
}

func (h *UserHandler) LoginUser(c *gin.Context) {

	creds := &entity.DefaultCredentials{}

	if err := c.Bind(creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.UserService.Login(creds)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User logged successfully :D", "token": token})

}
