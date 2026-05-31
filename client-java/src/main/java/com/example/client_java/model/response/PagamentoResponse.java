package com.example.client_java.model.response;

public record PagamentoResponse(
        String mensagem,
        int numeroConta,
        String descricao,
        double valorPago,
        double saldoAtual
    ) {}
