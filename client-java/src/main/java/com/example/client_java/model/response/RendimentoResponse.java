package com.example.client_java.model.response;

public record RendimentoResponse(
        int numeroConta,
        double saldoAtual,
        int meses,
        double taxaMensal,
        double saldoProjetado,
        double rendimento
    ) {}
