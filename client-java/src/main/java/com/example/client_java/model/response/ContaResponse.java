package com.example.client_java.model.response;

public record ContaResponse(
        int numero,
        String titular,
        double saldo,
        int tipo,
        double limite,
        double taxa
    ) {
}
