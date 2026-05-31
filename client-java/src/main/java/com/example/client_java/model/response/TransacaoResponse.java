package com.example.client_java.model.response;

public record TransacaoResponse(
        String tipo,
        double valor,
        String descricao,
        String data
    ) {}
