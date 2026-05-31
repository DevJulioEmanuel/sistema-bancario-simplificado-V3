package com.example.client_java.model.response;

import java.util.List;

public record ExtratoResponse(
        int numeroConta,
        String titular,
        double saldoAtual,
        List<TransacaoResponse> transacoes
    ) {}
