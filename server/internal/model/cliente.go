package model

type Cliente struct {
	Nome  string `json:"nome"`
	CPF   string `json:"cpf"`
	Senha string `json:"-"`
}
