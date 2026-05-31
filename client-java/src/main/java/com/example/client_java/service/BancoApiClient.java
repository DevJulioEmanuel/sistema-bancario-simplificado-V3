package com.example.client_java.service;

import com.example.client_java.model.request.*;
import com.example.client_java.model.response.*;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

@Service
public class BancoApiClient {

    private final RestClient restClient;

    public BancoApiClient(RestClient restClient) {
        this.restClient = restClient;
    }

    public CadastroResponse cadastrar(CadastroRequest cadastroRequest) {
        return restClient.post()
                .uri("/clientes")
                .body(cadastroRequest)
                .retrieve()
                .body(CadastroResponse.class);
    }

    public LoginResponse login(LoginRequest loginRequest) {
        return restClient.post()
                .uri("/clientes/login")
                .body(loginRequest)
                .retrieve()
                .body(LoginResponse.class);
    }

    public ContaResponse obterDados(int numero) {
        return restClient.get()
                .uri("/contas/" + numero + "/")
                .retrieve()
                .body(ContaResponse.class);
    }

    public PagamentoResponse pagar(int numero, PagamentoRequest  pagamentoRequest) {
        return restClient.post()
                .uri("/contas/" + numero + "/pagamento")
                .body(pagamentoRequest)
                .retrieve()
                .body(PagamentoResponse.class);
    }

    public MovimentacaoResponse depositar(int numero, ValorRequest  valorRequest) {
        return restClient.post()
                .uri("/contas/" + numero + "/depositar")
                .body(valorRequest)
                .retrieve()
                .body(MovimentacaoResponse.class);
    }

    public MovimentacaoResponse sacar(int numero, ValorRequest  valorRequest) {
        return restClient.post()
                .uri("/contas/" + numero + "/sacar")
                .body(valorRequest)
                .retrieve()
                .body(MovimentacaoResponse.class);
    }

    public TransferenciaResponse transferir(int numero, TransferenciaRequest valorRequest) {
        return restClient.post()
                .uri("/contas/" + numero + "/transferir")
                .body(valorRequest)
                .retrieve()
                .body(TransferenciaResponse.class);
    }

    public ExtratoResponse extrato(int numero) {
        return restClient.get()
                .uri("/contas/" + numero + "/extrato")
                .retrieve()
                .body(ExtratoResponse.class);
    }

    public RendimentoResponse rendimento(int numero, RendimentoRequest rendimentoRequest) {
        return restClient.get()
                .uri("/contas/" + numero + "/rendimento/" + rendimentoRequest.meses())
                .retrieve()
                .body(RendimentoResponse.class);
    }

}
