package request

type LoginRequest struct {
	CPF   string `json:"cpf" binding:"required"`
	Senha string `json:"senha" binding:"required"`
}
