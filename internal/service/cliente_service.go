package service

import (
	"banco-api/internal/model"
	"banco-api/internal/repository"
	"errors"
)

type ClienteService struct {
	clienteRepo *repository.ClienteRepository
	contaRepo   *repository.ContaRepository
}

func NewClienteService(
	clienteRepo *repository.ClienteRepository,
	contaRepo *repository.ContaRepository,
) *ClienteService {
	return &ClienteService{
		clienteRepo: clienteRepo,
		contaRepo:   contaRepo,
	}
}

func (s *ClienteService) Cadastrar(
	nome,
	cpf,
	senha string,
	tipo model.TipoConta,
) (*model.Conta, error) {

	_, existe := s.clienteRepo.BuscarPorCPF(cpf)

	if existe {
		return nil, errors.New("CPF já cadastrado")
	}

	cliente := &model.Cliente{
		Nome:  nome,
		CPF:   cpf,
		Senha: senha,
	}

	s.clienteRepo.Salvar(cliente)

	conta := &model.Conta{
		Numero:  s.contaRepo.GerarNumeroConta(tipo),
		Saldo:   0,
		Titular: cliente,
		Tipo:    tipo,
	}

	s.contaRepo.Salvar(conta)

	return conta, nil
}

func (s *ClienteService) Login(cpf, senha string) (*model.Cliente, error) {

	cliente, existe := s.clienteRepo.BuscarPorCPF(cpf)

	if !existe {
		return nil, errors.New("cliente não encontrado")
	}

	if cliente.Senha != senha {
		return nil, errors.New("senha incorreta")
	}

	return cliente, nil
}
