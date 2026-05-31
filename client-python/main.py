import time
import sys
import sessao
from ui.banner import Banner
from service.api_client import BancoApiClient
from ui.tela_login import executar_login
from ui.tela_cadastro import executar_cadastro
from ui.tela_conta import exibir_tela_conta


def main():
    api_client = BancoApiClient()

    Banner.limpar()
    Banner.exibir_cabecalho()
    _pausa(1000)

    rodando = True
    while rodando:
        Banner.limpar()
        Banner.exibir_cabecalho("MENU PRINCIPAL")
        _exibir_menu_principal()

        opcao = input().strip()

        if opcao == "1":
            sucesso = executar_login(api_client)
            if sucesso and sessao.esta_logado():
                exibir_tela_conta(api_client)

        elif opcao == "2":
            executar_cadastro(api_client)

        elif opcao == "0":
            Banner.limpar()
            Banner.exibir_cabecalho()
            Banner.info("Obrigado por usar o Banco API. Até logo!")
            _pausa(1500)
            rodando = False

        else:
            Banner.erro("Opção inválida. Tente novamente.")
            _pausa(1000)

    sys.exit(0)


def _exibir_menu_principal():
    m = Banner.margem()
    separador = getattr(Banner, 'SEPARADOR', "━" * 50)

    print(f"{m}{separador}")
    print(f"{m}  [1]  Entrar na minha conta")
    print(f"{m}  [2]  Abrir nova conta")
    print(f"{m}  [0]  Sair")
    print(f"{m}{separador}")

    Banner.prompt("Escolha:")


def _pausa(milissegundos: int):
    time.sleep(milissegundos / 1000.0)


if __name__ == "__main__":
    main()