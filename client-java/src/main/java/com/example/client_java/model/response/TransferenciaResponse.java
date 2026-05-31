package com.example.client_java.model.response;

public record TransferenciaResponse(
        String mensagem,
        int contaOrigem,
        int contaDestino,
        double valorTransferido,
        double saldoAtualOrigem
    ) {}
