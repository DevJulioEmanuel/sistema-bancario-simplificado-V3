package handler

import (
	dtorequest "banco-api/internal/dto/request"
	dtoresponse "banco-api/internal/dto/response"
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

func (h *ContaHandler) ObterDados(c *gin.Context) {
	numero, _ := strconv.Atoi(c.Param("num"))

	conta, err := h.service.BuscarConta(numero)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	tipo := 1

	if conta.Tipo == 2 {
		tipo = 2
	}

	response := dtoresponse.ContaResponse{
		Numero:  conta.Numero,
		Titular: conta.Titular.Nome,
		Saldo:   conta.Saldo,
		Tipo:    tipo,
		Limite:  1200,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ContaHandler) Pagar(c *gin.Context) {

	numero, err := strconv.Atoi(c.Param("num"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "número inválido",
		})
		return
	}

	var req dtorequest.PagamentoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	conta, err := h.service.Pagar(
		numero,
		req.Descricao,
		req.Valor,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	response := dtoresponse.PagamentoResponse{
		Mensagem:    "pagamento realizado com sucesso",
		NumeroConta: conta.Numero,
		Descricao:   req.Descricao,
		ValorPago:   req.Valor,
		SaldoAtual:  conta.Saldo,
	}

	c.JSON(http.StatusOK, response)
}
func (h *ContaHandler) Depositar(c *gin.Context) {
	numero, _ := strconv.Atoi(c.Param("num"))

	var req dtorequest.ValorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	conta, err := h.service.Depositar(numero, req.Valor)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	response := dtoresponse.MovimentacaoResponse{
		Mensagem:         "depósito realizado com sucesso",
		NumeroConta:      conta.Numero,
		TipoOperacao:     "deposito",
		ValorMovimentado: req.Valor,
		SaldoAtual:       conta.Saldo,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ContaHandler) Sacar(c *gin.Context) {
	numero, _ := strconv.Atoi(c.Param("num"))

	var req dtorequest.ValorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	conta, err := h.service.Sacar(numero, req.Valor)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	response := dtoresponse.MovimentacaoResponse{
		Mensagem:         "saque realizado com sucesso",
		NumeroConta:      conta.Numero,
		TipoOperacao:     "saque",
		ValorMovimentado: req.Valor,
		SaldoAtual:       conta.Saldo,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ContaHandler) Transferir(c *gin.Context) {
	numero, _ := strconv.Atoi(c.Param("num"))

	var req dtorequest.TransferenciaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	conta, err := h.service.Transferir(numero, req.NumDestino, req.Valor)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	response := dtoresponse.TransferenciaResponse{
		Mensagem:         "transferência realizada com sucesso",
		ContaOrigem:      conta.Numero,
		ContaDestino:     req.NumDestino,
		ValorTransferido: req.Valor,
		SaldoAtualOrigem: conta.Saldo,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ContaHandler) Extrato(c *gin.Context) {
	numero, err := strconv.Atoi(c.Param("num"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "número inválido",
		})
		return
	}

	conta, err := h.service.ObterExtrato(numero)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var transacoes []dtoresponse.TransacaoResponse

	for _, t := range conta.Historico {
		transacoes = append(transacoes, dtoresponse.TransacaoResponse{
			Tipo:      t.Tipo,
			Descricao: t.Descricao,
			Valor:     t.Valor,
			Data:      t.Data.Format("02/01/2006 15:04:05"),
		})
	}

	response := dtoresponse.ExtratoResponse{
		NumeroConta: conta.Numero,
		Titular:     conta.Titular.Nome,
		SaldoAtual:  conta.Saldo,
		Transacoes:  transacoes,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ContaHandler) CalcularRendimento(c *gin.Context) {

	numero, err := strconv.Atoi(c.Param("num"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "número inválido",
		})
		return
	}

	meses, err := strconv.Atoi(c.Param("meses"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "quantidade de meses inválida",
		})
		return
	}

	saldoProjetado, rendimento, conta, err :=
		h.service.CalcularRendimento(numero, meses)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	response := dtoresponse.RendimentoResponse{
		NumeroConta:    conta.Numero,
		SaldoAtual:     conta.Saldo,
		Meses:          meses,
		TaxaMensal:     0.005,
		SaldoProjetado: saldoProjetado,
		Rendimento:     rendimento,
	}

	c.JSON(http.StatusOK, response)
}
