package repository

import (
	"banco-api/internal/model"
	"sync"
)

type ContaRepository struct {
	contas          map[int]*model.Conta
	proximoCorrente int
	proximoPoupanca int
	mu              sync.RWMutex
}

func NewContaRepository() *ContaRepository {
	return &ContaRepository{
		contas:          make(map[int]*model.Conta),
		proximoCorrente: 1001,
		proximoPoupanca: 5001,
	}
}

func (r *ContaRepository) Salvar(conta *model.Conta) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.contas[conta.Numero] = conta
}

func (r *ContaRepository) BuscarPorNumero(numero int) (*model.Conta, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	conta, ok := r.contas[numero]

	return conta, ok
}

func (r *ContaRepository) GerarNumeroConta(tipo model.TipoConta) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	if tipo == 1 {
		num := r.proximoCorrente
		r.proximoCorrente++
		return num
	} else {
		num := r.proximoPoupanca
		r.proximoPoupanca++
		return num
	}
}
