package com.example.client_java.model.response;

public record MovimentacaoResponse(
        String mensagem,
        int numeroConta,
        String tipoOperacao,
        double valorMovimentado,
        double saldoAtual
    ) {}
