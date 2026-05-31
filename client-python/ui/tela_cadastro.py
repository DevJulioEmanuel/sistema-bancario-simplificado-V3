# ui/tela_cadastro.py
import time
import getpass
import sys
from ui.banner import Banner
from models import CadastroRequest

TIPOS = [
    ["1", "Corrente", "Limite R$ 1.200,00", "CORRENTE"],
    ["2", "Poupança", "Rendimento 0,500% a.m.", "POUPANÇA"],
]


def executar_cadastro(api_client):
    Banner.limpar()
    Banner.exibir_cabecalho("ABERTURA DE CONTA")
    Banner.titulo_secao("DADOS PESSOAIS")

    nome = _ler_obrigatorio("Nome completo  :")
    if nome is None: return

    cpf = _ler_obrigatorio("CPF            :")
    if cpf is None: return

    senha = _ler_senha("Senha          :")
    if senha is None: return

    confirma_senha = _ler_senha("Confirme senha :")
    if confirma_senha is None: return

    if senha != confirma_senha:
        Banner.erro("As senhas não coincidem.")
        _pausa(2000)
        return

    Banner.espaco()
    Banner.label_secao("TIPO DE CONTA")
    Banner.espaco()

    for tipo_item in TIPOS:
        detalhe = Banner.cinza(f"({tipo_item[2]})")
        print(f"{Banner.margem()}  "
              f"{Banner.cinza('[' + tipo_item[0] + ']  ')}"
              f"{Banner.branco(f'{tipo_item[1]:<12}')}  {detalhe}")

    print(f"{Banner.margem()}  {Banner.cinza('[0]  Voltar')}")
    Banner.espaco()
    Banner.prompt("Tipo            :")
    tipo_str = input().strip()

    if tipo_str == "0": return

    try:
        tipo = int(tipo_str)
        if tipo < 1 or tipo > 2:
            raise ValueError()
    except ValueError:
        Banner.erro("Tipo inválido.")
        _pausa(1500)
        return

    Banner.espaco()
    Banner.linha_separadora()
    Banner.espaco()
    _exibir_resumo(nome, cpf, tipo)
    Banner.espaco()
    Banner.linha_separadora()
    Banner.espaco()

    print(f"{Banner.margem()}  {Banner.cinza('[1]  Confirmar abertura')}")
    print(f"{Banner.margem()}  {Banner.cinza('[0]  Cancelar')}")
    Banner.espaco()
    Banner.prompt("Escolha        :")
    conf = input().strip()

    if conf != "1":
        Banner.info("Operação cancelada.")
        _pausa(1200)
        return

    Banner.espaco()
    Banner.info("Cadastrando...")
    _pausa(500)

    try:
        request = CadastroRequest(nome=nome, cpf=cpf, senha=senha, tipoConta=tipo)

        api_client.cadastrar(request)

        Banner.sucesso("Conta criada com sucesso!")
        Banner.info("Você já pode fazer login com seu CPF e senha.")
        _pausa(2500)

    except Exception as e:
        Banner.erro(f"Erro ao cadastrar: {e}")
        _pausa(2000)

def _exibir_resumo(nome: str, cpf: str, tipo: int):
    info = TIPOS[tipo - 1]
    Banner.label_secao("RESUMO DA ABERTURA")
    Banner.campo("Titular", Banner.branco(nome))
    Banner.campo("CPF", Banner.cinza(cpf))

    badge_tipo = Banner.badge_azul(info[3]) if tipo == 1 else Banner.badge_verde(info[3])
    Banner.campo("Tipo", badge_tipo)
    Banner.campo("Condições", Banner.cinza(info[2]))


def _ler_obrigatorio(label: str) -> str or None:
    Banner.prompt(label)
    valor = input().strip()
    if not valor:
        Banner.erro("Campo obrigatório.")
        _pausa(1000)
        return None
    return valor


def _ler_senha(label: str) -> str or None:
    print(f"{Banner.margem()}  {Banner.B_AZUL}{label}{Banner.RESET} ", end="")
    sys.stdout.flush()

    valor = getpass.getpass("").strip()

    if not valor:
        Banner.erro("Senha obrigatória.")
        _pausa(1000)
        return None
    return valor


def _pausa(milissegundos: int):
    time.sleep(milissegundos / 1000.0)