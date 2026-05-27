package model

import (
	"sync"
)

type TipoConta int

const (
	TipoCorrente TipoConta = 1
	TipoPoupanca TipoConta = 2
	TaxaMensal             = 0.005
)

type Conta struct {
	Numero    int
	Saldo     float64
	Titular   *Cliente
	Tipo      TipoConta
	Historico []Transacao
	Mu        sync.RWMutex
}
