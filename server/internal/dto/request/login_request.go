package request

import "banco-api/internal/model"

type LoginRequest struct {
	CPF       string          `json:"cpf" binding:"required"`
	Senha     string          `json:"senha" binding:"required"`
	TipoConta model.TipoConta `json:"tipoConta"`
}
