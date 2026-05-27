package routes

import (
	"banco-api/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	clienteHandler *handler.ClienteHandler,
	contaHandler *handler.ContaHandler,
) {

	clientes := r.Group("/clientes")
	{
		clientes.POST("", clienteHandler.Cadastrar)
		clientes.POST("/login", clienteHandler.Login)
	}

	contas := r.Group("/contas/:num")
	{
		contas.POST("/depositar", contaHandler.Depositar)
		contas.POST("/sacar", contaHandler.Sacar)
		contas.POST("/transferir", contaHandler.Transferir)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
