package com.example.client_java.model.request;

public record LoginRequest(
        String cpf,
        String senha,
        int tipoConta
    ) {
}
