package com.example.client_java.ui;

import com.example.client_java.model.ContaLogada;
import com.example.client_java.model.request.LoginRequest;
import com.example.client_java.model.response.ContaResponse;
import com.example.client_java.model.response.LoginResponse;
import com.example.client_java.service.BancoApiClient;
import org.springframework.stereotype.Component;
import org.springframework.web.client.HttpClientErrorException;

import java.util.Scanner;

@Component
public class TelaLogin {
    private final Scanner sc;
    private final BancoApiClient bancoApiClient;

    public TelaLogin(BancoApiClient bancoApiClient, Scanner sc) {
        this.bancoApiClient = bancoApiClient;
        this.sc = sc;
    }

    public ContaLogada exibir() {
        Banner.limpar();
        Banner.exibirCabecalho("ACESSO À CONTA");
        Banner.tituloSecao("IDENTIFICAÇÃO DO CLIENTE");

        System.out.print("CPF          : ");
        String cpf = sc.nextLine().trim();

        System.out.print("Senha        : ");
        String senha = sc.nextLine().trim();

        System.out.print("Tipo da Conta [1-Corrente, 2-Poupança]: ");
        int tipo = Integer.parseInt(sc.nextLine().trim());

        Banner.espaco();
        Banner.linhaSeparadora();
        Banner.info("Autenticando...");
        pausa(500);

        try {
            LoginRequest req = new LoginRequest(cpf, senha, tipo);

            LoginResponse resposta = bancoApiClient.login(req);

            ContaResponse response = bancoApiClient.obterDados(resposta.numero());

            Banner.sucesso(resposta.mensagem());
            pausa(1000);

            return new ContaLogada(
                    response.numero(),
                    cpf,
                    resposta.nome(),
                    tipo,
                    response.saldo(),
                    response.limite(),
                    response.taxa()
            );

        } catch (HttpClientErrorException e) {
            String erro = e.getResponseBodyAsString();
            String mensagem = erro
                    .replaceAll(".*\"erro\":\"*", "")
                    .replaceAll("\"}.*", "")
                    .trim();

            Banner.erro("Falha no login: " + mensagem);
            pausa(2000);

            return null;
        }
    }

    private void pausa(long ms) {
        try { Thread.sleep(ms); } catch (InterruptedException ignored) {}
    }
}
