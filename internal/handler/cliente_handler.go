package handler

import (
	"banco-api/internal/dto"
	"banco-api/internal/model"
	"banco-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	service *service.ClienteService
}

func NewClienteHandler(service *service.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: service}
}

func (h *ClienteHandler) Cadastrar(c *gin.Context) {
	var req dto.CadastroRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	conta, err := h.service.Cadastrar(
		req.Nome,
		req.CPF,
		req.Senha,
		model.TipoConta(req.Tipo),
	)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensagem":  "cliente cadastrado com sucesso",
		"num_conta": conta.Numero,
	})
}

func (h *ClienteHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	cliente, err := h.service.Login(req.CPF, req.Senha)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "login realizado",
		"nome":     cliente.Nome,
	})
}
