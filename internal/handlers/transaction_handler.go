package handlers

import (
	"dummy-payment-processing/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {

	var req dto.CreateTransactionRequest
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.service.CreateTransaction(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (h *TransactionHandler) GetTransactionStatus(c *gin.Context) {

	id := c.Param("id")
	ctx := c.Request.Context()

	resp, err := h.service.GetTransactionStatus(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (h *TransactionHandler) GetTransactionStats(c *gin.Context) {

	ctx := c.Request.Context()

	resp, err := h.service.GetTransactionStats(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {

	ctx := c.Request.Context()

	resp := h.service.GetAllTransactions(ctx)

	c.JSON(http.StatusOK, resp)

}
