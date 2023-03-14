package deposit

import (
	"net/http"

	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/entity"
	"github.com/OscarLlamas6/hex-arch-go/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type (
	DepositHandler struct {
		DepositService ports.DepositService
	}
)

func NewDepositHandler(depositService ports.DepositService) *DepositHandler {

	return &DepositHandler{
		DepositService: depositService,
	}
}

func (h *DepositHandler) CreateDeposit(c *gin.Context) {
	deposit := &entity.Deposit{}

	if err := c.Bind(deposit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DepositService.AddDeposit(deposit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Deposit created successfully :D"})
}
