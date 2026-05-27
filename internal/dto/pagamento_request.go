package dto

type PagamentoRequest struct {
	Valor     float64 `json:"valor" binding:"required,gt=0"`
	Descricao string  `json:"descricao" binding:"required"`
}
