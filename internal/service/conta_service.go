package service

import (
	"banco-api/internal/model"
	"banco-api/internal/repository"
	"errors"
	"fmt"
	"time"
)

type ContaService struct {
	repo *repository.ContaRepository
}

func NewContaService(repo *repository.ContaRepository) *ContaService {
	return &ContaService{repo: repo}
}

func (s *ContaService) BuscarConta(numero int) (*model.Conta, error) {

	conta, ok := s.repo.BuscarPorNumero(numero)

	if !ok {
		return nil, errors.New("conta não encontrada")
	}

	return conta, nil
}

func (s *ContaService) Depositar(numero int, valor float64) error {

	conta, err := s.BuscarConta(numero)

	if err != nil {
		return err
	}

	if valor <= 0 {
		return errors.New("valor inválido")
	}

	conta.Mu.Lock()
	defer conta.Mu.Unlock()

	conta.Saldo += valor

	conta.Historico = append(conta.Historico, model.Transacao{
		Descricao: "Depósito",
		Valor:     valor,
		Data:      time.Now(),
	})

	return nil
}

func (s *ContaService) Sacar(numero int, valor float64) error {

	conta, err := s.BuscarConta(numero)

	if err != nil {
		return err
	}

	if valor <= 0 {
		return errors.New("valor inválido")
	}

	conta.Mu.Lock()
	defer conta.Mu.Unlock()

	if conta.Saldo < valor {
		return errors.New("saldo insuficiente")
	}

	conta.Saldo -= valor

	conta.Historico = append(conta.Historico, model.Transacao{
		Descricao: "Saque",
		Valor:     -valor,
		Data:      time.Now(),
	})

	return nil
}

func (s *ContaService) Transferir(origemNum, destinoNum int, valor float64) error {

	origem, err := s.BuscarConta(origemNum)

	if err != nil {
		return err
	}

	destino, err := s.BuscarConta(destinoNum)

	if err != nil {
		return errors.New("conta destino não encontrada")
	}

	if valor <= 0 {
		return errors.New("valor inválido")
	}

	primeiro := origem
	segundo := destino

	if origem.Numero > destino.Numero {
		primeiro = destino
		segundo = origem
	}

	primeiro.Mu.Lock()
	defer primeiro.Mu.Unlock()

	segundo.Mu.Lock()
	defer segundo.Mu.Unlock()

	if origem.Saldo < valor {
		return errors.New("saldo insuficiente")
	}

	origem.Saldo -= valor
	destino.Saldo += valor

	origem.Historico = append(origem.Historico, model.Transacao{
		Descricao: fmt.Sprintf("Transferência para conta %d", destino.Numero),
		Valor:     -valor,
		Data:      time.Now(),
	})

	destino.Historico = append(destino.Historico, model.Transacao{
		Descricao: fmt.Sprintf("Transferência recebida da conta %d", origem.Numero),
		Valor:     valor,
		Data:      time.Now(),
	})

	return nil
}
