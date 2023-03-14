package deposit

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/api/middlewares"
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/repositories/couchdb"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/service/deposit"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo, err := couchdb.NewCouchDBClient()
	if err != nil {
		panic(err)
	}

	service := deposit.NewService(repo)
	handler := NewDepositHandler(service)

	e.POST("/api/v1/deposit", middlewares.Authenticate(), handler.CreateDeposit)
}
