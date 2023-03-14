package expense

import (
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/api/middlewares"
	"github.com/OscarLlamas6/hex-arch-go/internal/infra/repositories/couchdb"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/service/expense"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo, err := couchdb.NewCouchDBClient()
	if err != nil {
		panic(err)
	}

	service := expense.NewService(repo)
	handler := NewExpenseHandler(service)

	e.POST("/api/v1/expense", middlewares.Authenticate(), handler.CreateExpense)
}
