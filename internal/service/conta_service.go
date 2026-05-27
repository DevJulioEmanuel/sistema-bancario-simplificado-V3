package service

import (
	"banco-api/internal/model"
	"banco-api/internal/repository"
	"errors"
	"fmt"
	"math"
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

func (s *ContaService) Pagar(
	numero int,
	descricao string,
	valor float64,
) (*model.Conta, error) {

	conta, err := s.BuscarConta(numero)

	if err != nil {
		return nil, err
	}

	if conta.Tipo != model.TipoCorrente {
		return nil, errors.New(
			"apenas contas correntes podem realizar pagamentos",
		)
	}

	if valor <= 0 {
		return nil, errors.New("valor inválido")
	}

	if conta.Saldo < valor {
		return nil, errors.New("saldo insuficiente")
	}

	conta.Mu.Lock()
	defer conta.Mu.Unlock()

	conta.Saldo -= valor

	transacao := model.Transacao{
		Tipo:      "pagamento",
		Descricao: descricao,
		Valor:     -valor,
		Data:      time.Now(),
	}

	conta.Historico = append(conta.Historico, transacao)

	return conta, nil
}

func (s *ContaService) Depositar(numero int, valor float64) (*model.Conta, error) {

	conta, err := s.BuscarConta(numero)

	if err != nil {
		return nil, err
	}

	if valor <= 0 {
		return nil, errors.New("valor inválido")
	}

	conta.Mu.Lock()
	defer conta.Mu.Unlock()

	conta.Saldo += valor

	conta.Historico = append(conta.Historico, model.Transacao{
		Tipo:      "deposito",
		Descricao: "Depósito",
		Valor:     valor,
		Data:      time.Now(),
	})

	return conta, nil
}

func (s *ContaService) Sacar(numero int, valor float64) (*model.Conta, error) {

	conta, err := s.BuscarConta(numero)

	if err != nil {
		return nil, err
	}

	if valor <= 0 {
		return nil, errors.New("valor inválido")
	}

	conta.Mu.Lock()
	defer conta.Mu.Unlock()

	if conta.Saldo < valor {
		return nil, errors.New("saldo insuficiente")
	}

	conta.Saldo -= valor

	conta.Historico = append(conta.Historico, model.Transacao{
		Tipo:      "saque",
		Descricao: "Saque",
		Valor:     -valor,
		Data:      time.Now(),
	})

	return conta, nil
}

func (s *ContaService) Transferir(origemNum, destinoNum int, valor float64) (*model.Conta, error) {

	origem, err := s.BuscarConta(origemNum)

	if err != nil {
		return nil, err
	}

	destino, err := s.BuscarConta(destinoNum)

	if err != nil {
		return nil, errors.New("conta destino não encontrada")
	}

	if valor <= 0 {
		return nil, errors.New("valor inválido")
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
		return nil, errors.New("saldo insuficiente")
	}

	origem.Saldo -= valor
	destino.Saldo += valor

	origem.Historico = append(origem.Historico, model.Transacao{
		Tipo:      "transferencia",
		Descricao: fmt.Sprintf("Transferência para conta %d", destino.Numero),
		Valor:     -valor,
		Data:      time.Now(),
	})

	destino.Historico = append(destino.Historico, model.Transacao{
		Tipo:      "transferencia",
		Descricao: fmt.Sprintf("Transferência recebida da conta %d", origem.Numero),
		Valor:     valor,
		Data:      time.Now(),
	})

	return origem, nil
}

func (s *ContaService) ObterExtrato(numero int) (*model.Conta, error) {
	conta, err := s.BuscarConta(numero)

	if err != nil {
		return nil, err
	}

	return conta, nil
}

func (s *ContaService) CalcularRendimento(numero int, meses int) (float64, float64, *model.Conta, error) {

	conta, err := s.BuscarConta(numero)

	if err != nil {
		return 0, 0, nil, err
	}

	if conta.Tipo != model.TipoPoupanca {
		return 0, 0, nil, errors.New("apenas contas poupanca possuem rendimento")
	}

	saldoProjetado := conta.Saldo * math.Pow(1+model.TaxaMensal, float64(meses))

	rendimento := saldoProjetado - conta.Saldo

	saldoProjetado = math.Round(saldoProjetado*100) / 100

	rendimento = math.Round(rendimento*100) / 100

	return saldoProjetado, rendimento, conta, nil
}
