package repository

import (
	"banco-api/internal/model"
	"sync"
)

type ClienteRepository struct {
	clientes map[string]*model.Cliente
	mu       sync.RWMutex
}

func NewClienteRepository() *ClienteRepository {
	return &ClienteRepository{
		clientes: make(map[string]*model.Cliente),
	}
}

func (r *ClienteRepository) Salvar(cliente *model.Cliente) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.clientes[cliente.CPF] = cliente
}

func (r *ClienteRepository) BuscarPorCPF(cpf string) (*model.Cliente, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cliente, ok := r.clientes[cpf]

	return cliente, ok
}

func (r *ContaRepository) BuscarPorClienteETipo(
	cpf string,
	tipo model.TipoConta,
) (*model.Conta, bool) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, conta := range r.contas {

		if conta.Titular.CPF == cpf &&
			conta.Tipo == tipo {

			return conta, true
		}
	}

	return nil, false
}
