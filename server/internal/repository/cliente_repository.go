package repository

import (
	"banco-api/internal/model"
	"sync"
)

type ClienteRepository struct {
	bancoRepo *BancoRepository
	mu        sync.RWMutex
}

func NewClienteRepository(
	bancoRepo *BancoRepository,
) *ClienteRepository {
	return &ClienteRepository{
		bancoRepo: bancoRepo,
	}
}

func (r *ClienteRepository) Salvar(cliente *model.Cliente) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.bancoRepo.GetBanco().Clientes =
		append(r.bancoRepo.GetBanco().Clientes, cliente)
}

func (r *ClienteRepository) BuscarPorCPF(cpf string) (*model.Cliente, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, cliente := range r.bancoRepo.GetBanco().Clientes {

		if cliente.CPF == cpf {
			return cliente, true
		}
	}

	return nil, false
}
