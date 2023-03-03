package user

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/repositories/pg"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo, err := pg.NewClient()
	if err != nil {
		panic(err)
	}

	service := user.NewService(repo)
	handler := NewUserHandler(service)

	e.POST("/api/v1/user", handler.CreateUser)
	e.POST("/api/v1/login", handler.LoginUser)
}
