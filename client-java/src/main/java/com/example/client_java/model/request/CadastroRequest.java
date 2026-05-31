package com.example.client_java.model.request;

public record CadastroRequest(
        String nome,
        String cpf,
        String senha,
        int tipoConta
    ) {}
