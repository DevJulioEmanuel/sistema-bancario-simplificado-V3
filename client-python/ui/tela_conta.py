import time
import sessao
from models import ValorRequest, TransferenciaRequest, PagamentoRequest, RendimentoRequest
from ui.banner import Banner

def exibir_tela_conta(api_client):
    logado = True

    while logado:
        _renderizar_dashboard()
        _exibir_menu()

        opcao = input().strip()

        if opcao == "1":
            _operacao_depositar(api_client)
        elif opcao == "2":
            _operacao_sacar(api_client)
        elif opcao == "3":
            _operacao_transferir(api_client)
        elif opcao == "4":
            _operacao_extrato(api_client)
        elif opcao == "5":
            if sessao.usuario_atual.tipo == 1:  # 1 - Corrente
                _operacao_pagar(api_client)
            elif sessao.usuario_atual.tipo == 2:  # 2 - Poupança
                _operacao_projetar_rendimento(api_client)
            else:
                _opcao_invalida()
        elif opcao == "0":
            Banner.limpar()
            Banner.exibir_cabecalho()
            Banner.info("Sessão encerrada. Até logo!")
            _pausa(1500)
            sessao.encerrar()
            logado = False
        else:
            _opcao_invalida()


def _renderizar_dashboard():
    u = sessao.usuario_atual
    Banner.limpar()
    Banner.exibir_cabecalho("MINHA CONTA")
    Banner.titulo_secao("PAINEL DO CLIENTE")

    Banner.label_secao("TITULAR")
    Banner.campo("Nome", Banner.branco(u.nome))
    Banner.campo("CPF", Banner.cinza(u.cpf))
    Banner.espaco()

    Banner.label_secao("CONTA")
    Banner.campo("Número", Banner.branco(str(u.numero)))
    Banner.campo("Tipo", Banner.badge_azul("CORRENTE") if u.tipo == 1 else Banner.badge_verde("POUPANÇA"))
    Banner.espaco()

    Banner.label_secao("SALDO DISPONÍVEL")
    print(f"{Banner.margem()}  {Banner.valor_positivo(u.saldo)}")

    if u.tipo == 1:
        print(f"{Banner.margem()}  {Banner.cinza(f'Limite: R$ {u.limite:.2f}')}")
    else:
        rendimento = getattr(u, 'rendimento', 0.005)
        print(f"{Banner.margem()}  {Banner.cinza(f'Rendimento: {rendimento * 100:.3f}% a.m.')}")

    Banner.espaco()
    Banner.linha_separadora()
    Banner.espaco()


def _exibir_menu():
    m = Banner.margem()
    print(f"{m}  {Banner.cinza('[1]  Depositar')}")
    print(f"{m}  {Banner.cinza('[2]  Sacar')}")
    print(f"{m}  {Banner.cinza('[3]  Transferir')}")
    print(f"{m}  {Banner.cinza('[4]  Ver Extrato')}")

    if sessao.usuario_atual.tipo == 1:
        print(f"{m}  {Banner.cinza('[5]  Pagar Boleto / Conta')}")
    else:
        print(f"{m}  {Banner.cinza('[5]  Projetar Rendimento')}")

    print(f"{m}  {Banner.cinza('[0]  Sair (Logout)')}")
    Banner.espaco()
    Banner.prompt("O que deseja fazer?")


def _operacao_depositar(api_client):
    valor = _ler_valor("Valor do depósito")
    if valor is None: return

    Banner.info("Processando...")
    try:
        valor_deposito = ValorRequest(valor)
        api_client.depositar(sessao.usuario_atual.numero, valor_deposito)
        sessao.usuario_atual.saldo += valor
        Banner.sucesso("Depósito realizado com sucesso!")
    except Exception as e:
        Banner.erro(f"Erro ao depositar: {e}")
        _pausa(2000)
    _pausa(2000)


def _operacao_sacar(api_client):
    valor = _ler_valor("Valor do saque")
    if valor is None: return

    Banner.info("Processando...")
    try:
        valor_saque = ValorRequest(valor)
        api_client.sacar(sessao.usuario_atual.numero, valor_saque)
        sessao.usuario_atual.saldo -= valor
        Banner.sucesso("Saque realizado com sucesso!")
    except Exception as e:
        Banner.erro(f"Erro ao sacar: {e}")
        _pausa(2000)
    _pausa(2000)


def _operacao_transferir(api_client):
    Banner.espaco()
    Banner.label_secao("TRANSFERÊNCIA")

    Banner.prompt("Conta destino      :")
    dest_str = input().strip()
    try:
        num_destino = int(dest_str)
    except ValueError:
        Banner.erro("Número de conta inválido.")
        _pausa(1500)
        return

    if num_destino == sessao.usuario_atual.numero:
        Banner.erro("Não é possível transferir para a própria conta.")
        _pausa(2000)
        return

    valor = _ler_valor("Valor da transferência")
    if valor is None: return

    Banner.info("Processando...")
    try:
        valor_transferencia = TransferenciaRequest(num_destino, valor)
        api_client.transferir(sessao.usuario_atual.numero, valor_transferencia)
        sessao.usuario_atual.saldo -= valor
        Banner.sucesso("Transferência realizada com sucesso!")
    except Exception as e:
        Banner.erro(f"Erro ao transferir: {e}")
        _pausa(2000)
    _pausa(2500)


def _operacao_pagar(api_client):
    Banner.espaco()
    Banner.label_secao("PAGAMENTO DE BOLETO / CONTA")

    Banner.prompt("Descrição          :")
    descricao = input().strip()
    if not descricao:
        Banner.erro("Descrição obrigatória.")
        _pausa(1500)
        return

    valor = _ler_valor("Valor do pagamento")
    if valor is None: return

    Banner.info("Processando...")
    try:
        valor_pagamento = PagamentoRequest(descricao, valor)
        api_client.pagar(sessao.usuario_atual.numero, valor_pagamento)

        sessao.usuario_atual.saldo -= valor
        Banner.sucesso(f"Pagamento efetuado com sucesso! — {descricao}")
    except Exception as e:
        Banner.erro(f"Erro ao pagar: {e}")
        _pausa(2000)
    _pausa(2000)


def _operacao_projetar_rendimento(api_client):
    Banner.espaco()
    Banner.label_secao("PROJEÇÃO DE RENDIMENTO")
    Banner.prompt("Número de meses    :")
    meses_str = input().strip()
    try:
        meses = int(meses_str)
        if meses <= 0: raise ValueError()
    except ValueError:
        Banner.erro("Número de meses inválido.")
        _pausa(1500)
        return

    Banner.info("Calculando...")
    try:
        valor_meses = RendimentoRequest(meses)
        res = api_client.rendimento(sessao.usuario_atual.numero, valor_meses)
        Banner.espaco()
        Banner.label_secao("RESULTADO DA PROJEÇÃO")
        Banner.campo("Período", f"{meses} meses")
        Banner.campo("Saldo atual", Banner.formatar_dinheiro(sessao.usuario_atual.saldo))

        Banner.campo("Saldo projetado", Banner.valor_positivo(res.saldoProjetado))
        Banner.espaco()
        Banner.aguarde_enter()
    except Exception as e:
        Banner.erro(f"Erro ao projetar rendimento: {e}")
        _pausa(1000000)


def _operacao_extrato(api_client):
    Banner.limpar()
    Banner.exibir_cabecalho("EXTRATO")
    Banner.titulo_secao(f"MOVIMENTAÇÕES DA CONTA {sessao.usuario_atual.numero}")

    Banner.info("Buscando movimentações...")
    try:
        res = api_client.extrato(sessao.usuario_atual.numero)
        Banner.espaco()

        lista_transacoes = res.transacoes if hasattr(res, 'transacoes') else getattr(res, 'transacoes', [])

        if not lista_transacoes:
            Banner.info("Nenhuma movimentação registrada.")
        else:
            Banner.label_secao("HISTÓRICO")
            for t in lista_transacoes:
                valor = t.valor if hasattr(t, 'valor') else t.get("valor", 0)
                desc = t.descricao if hasattr(t, 'descricao') else t.get("descricao", "")

                if valor >= 0:
                    valor_formatado = f"{Banner.B_VERDE}+R$ {valor:,.2f}{Banner.RESET}"
                else:
                    valor_formatado = f"{Banner.B_VERM}-R$ {abs(valor):,.2f}{Banner.RESET}"

                print(f"{Banner.margem()}  {Banner.cinza('• ')}{desc}: {valor_formatado}")
    except Exception as e:
        Banner.erro(f"Erro ao buscar extrato: {e}")
        _pausa(2000)

    Banner.espaco()
    Banner.linha_separadora()
    Banner.espaco()
    Banner.aguarde_enter()


def _ler_valor(label: str) -> float or None:
    Banner.espaco()
    Banner.prompt(f"{label}  : R$ ")
    s = input().strip().replace(",", ".")
    try:
        v = float(s)
        if v <= 0: raise ValueError()
        return v
    except ValueError:
        Banner.erro("Valor inválido. Informe um número positivo.")
        _pausa(1500)
        return None


def _opcao_invalida():
    Banner.erro("Opção inválida.")
    _pausa(1000)


def _pausa(milissegundos: int):
    time.sleep(milissegundos / 1000.0)