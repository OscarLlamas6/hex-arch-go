package expense

import (
	"net/http"

	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type (
	ExpenseHandler struct {
		ExpenseService ports.ExpenseService
	}
)

func NewExpenseHandler(expenseService ports.ExpenseRepository) *ExpenseHandler {

	return &ExpenseHandler{
		ExpenseService: expenseService,
	}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	expense := &entity.Expense{}

	if err := c.Bind(expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ExpenseService.AddExpense(expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Expense created successfully :D"})
}
