import time
import sessao
from ui.banner import Banner
from models import LoginRequest, ContaLogada


def executar_login(api_client) -> bool:
    Banner.limpar()
    Banner.exibir_cabecalho("ACESSO À CONTA")
    Banner.titulo_secao("IDENTIFICAÇÃO DO CLIENTE")

    Banner.prompt("CPF          :")
    cpf = input().strip()

    Banner.prompt("Senha        :")
    senha = input().strip()

    Banner.prompt("Tipo da Conta [1-Corrente, 2-Poupança]:")
    try:
        tipo = int(input().strip())
    except ValueError:
        Banner.erro("Tipo de conta inválido. Digite 1 ou 2.")
        _pausa(2000)
        return False

    Banner.espaco()
    Banner.linha_separadora()
    Banner.info("Autenticando...")
    _pausa(500)

    try:
        req = LoginRequest(cpf=cpf, senha=senha, tipoConta=tipo)

        resposta_login = api_client.login(req)

        numero_conta = resposta_login.get("numero") if isinstance(resposta_login, dict) else resposta_login.numero
        response_conta = api_client.obterDados(numero_conta)

        msg_sucesso = resposta_login.get("mensagem") if isinstance(resposta_login, dict) else resposta_login.mensagem
        Banner.sucesso(msg_sucesso)
        _pausa(1000)

        nome_titular = resposta_login.get("nome") if isinstance(resposta_login, dict) else resposta_login.nome

        conta_logada = ContaLogada(
            numero=numero_conta,
            nome=nome_titular,
            cpf=cpf,
            tipo=tipo,
            saldo=response_conta.saldo,
            limite=response_conta.limite,
            rendimento=response_conta.taxa,
        )


        sessao.iniciar(conta_logada)
        return True

    except Exception as e:
        Banner.erro(f"Falha no login: {e}")
        _pausa(2000)
        return False


def _pausa(milissegundos: int):
    time.sleep(milissegundos / 1000.0)