package repository

import (
	"banco-api/internal/model"
	"sync"
)

type ContaRepository struct {
	contas        map[int]*model.Conta
	proximoNumero int
	mu            sync.RWMutex
}

func NewContaRepository() *ContaRepository {
	return &ContaRepository{
		contas:        make(map[int]*model.Conta),
		proximoNumero: 1001,
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

func (r *ContaRepository) GerarNumeroConta() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	numero := r.proximoNumero
	r.proximoNumero++

	return numero
}
