package dto

type TransferenciaRequest struct {
	NumDestino int     `json:"num_destino" binding:"required"`
	Valor      float64 `json:"valor" binding:"required,gt=0"`
}
