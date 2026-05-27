package dto

type CadastroRequest struct {
	Nome  string `json:"nome" binding:"required"`
	CPF   string `json:"cpf" binding:"required"`
	Senha string `json:"senha" binding:"required"`
	Tipo  int    `json:"tipo" binding:"required,oneof=1 2"`
}
