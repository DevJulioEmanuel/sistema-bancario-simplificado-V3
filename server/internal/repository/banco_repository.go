package repository

import "banco-api/internal/model"

type BancoRepository struct {
	banco *model.Banco
}

func NewBancoRepository() *BancoRepository {
	return &BancoRepository{
		banco: &model.Banco{
			Clientes: make([]*model.Cliente, 0),
			Contas:   make([]*model.Conta, 0),
		},
	}
}

func (r *BancoRepository) GetBanco() *model.Banco {
	return r.banco
}
