package repository

import (
	"banco-api/internal/model"
	"sync"
)

type ContaRepository struct {
	bancoRepo       *BancoRepository
	proximoCorrente int
	proximoPoupanca int
	mu              sync.RWMutex
}

func NewContaRepository(
	bancoRepo *BancoRepository,
) *ContaRepository {
	return &ContaRepository{
		bancoRepo:       bancoRepo,
		proximoCorrente: 1001,
		proximoPoupanca: 5001,
	}
}

func (r *ContaRepository) Salvar(conta *model.Conta) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.bancoRepo.GetBanco().Contas =
		append(r.bancoRepo.GetBanco().Contas, conta)
}

func (r *ContaRepository) BuscarPorNumero(numero int) (*model.Conta, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, conta := range r.bancoRepo.GetBanco().Contas {
		if conta.Numero == numero {
			return conta, true
		}
	}

	return nil, false
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

func (r *ContaRepository) BuscarPorClienteETipo(
	cpf string,
	tipo model.TipoConta,
) (*model.Conta, bool) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, conta := range r.bancoRepo.GetBanco().Contas {
		if conta.Titular.CPF == cpf &&
			conta.Tipo == tipo {

			return conta, true
		}
	}

	return nil, false
}
