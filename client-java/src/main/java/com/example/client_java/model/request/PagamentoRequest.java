package com.example.client_java.model.request;

public record PagamentoRequest(
        String descricao,
        double valor
    ) {}
