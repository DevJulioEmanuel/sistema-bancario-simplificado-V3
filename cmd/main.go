package main

import (
	"fmt"

	"banco-api/internal/handler"
	"banco-api/internal/repository"
	"banco-api/internal/routes"
	"banco-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	clienteRepo := repository.NewClienteRepository()
	contaRepo := repository.NewContaRepository()

	clienteService := service.NewClienteService(clienteRepo, contaRepo)
	contaService := service.NewContaService(contaRepo)

	clienteHandler := handler.NewClienteHandler(clienteService)
	contaHandler := handler.NewContaHandler(contaService)

	r := gin.Default()

	routes.SetupRoutes(r, clienteHandler, contaHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")

	r.Run(":8080")
}
