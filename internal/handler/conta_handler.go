package handler

import (
	"banco-api/internal/dto"
	"banco-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContaHandler struct {
	service *service.ContaService
}

func NewContaHandler(service *service.ContaService) *ContaHandler {
	return &ContaHandler{service: service}
}

func (h *ContaHandler) Depositar(c *gin.Context) {
	numero, _ := strconv.Atoi(c.Param("num"))

	var req dto.ValorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := h.service.Depositar(numero, req.Valor)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "depósito realizado",
	})
}

func (h *ContaHandler) Sacar(c *gin.Context) {
	numero, _ := strconv.Atoi(c.Param("num"))

	var req dto.ValorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := h.service.Sacar(numero, req.Valor)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "saque realizado",
	})
}

func (h *ContaHandler) Transferir(c *gin.Context) {
	numero, _ := strconv.Atoi(c.Param("num"))

	var req dto.TransferenciaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err := h.service.Transferir(numero, req.NumDestino, req.Valor)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "transferência realizada",
	})
}
