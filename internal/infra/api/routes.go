package api

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	user.RegisterRoutes(e)
}
