import os
import platform


class Banner:
    LARGURA = 60
    NOME = "BancoAPI"
    VERSAO = "v3.0 GO · Json · API"

    SEPARADOR = "─" * LARGURA
    SEPARADOR_DUPLO = "· " * (LARGURA // 2)

    # ── ANSI RESET ──────────────────────────────────────────────────
    RESET = "\u001B[0m"
    BOLD = "\u001B[1m"

    # ── Paleta para terminal claro ─────────────────────────────────
    CINZA = "\u001B[38;5;245m"
    AZUL = "\u001B[38;5;33m"
    VERDE = "\u001B[38;5;35m"
    VERMELHO = "\u001B[38;5;160m"
    AMARELO = "\u001B[38;5;136m"
    CIANO = "\u001B[38;5;37m"

    B_AZUL = BOLD + AZUL
    B_VERDE = BOLD + VERDE
    B_VERM = BOLD + VERMELHO
    B_AMAR = BOLD + AMARELO
    B_CIANO = BOLD + CIANO
    B_CINZA = BOLD + CINZA

    # ── Logo ────────────────────────────────────────────────────────
    LOGO = [
        " ██████╗  █████╗ ███╗   ██╗ ██████╗ ██████╗ ",
        " ██╔══██╗██╔══██╗████╗  ██║██╔════╝██╔═══██╗",
        " ██████╔╝███████║██╔██╗ ██║██║     ██║   ██║",
        " ██╔══██╗██╔══██║██║╚██╗██║██║     ██║   ██║",
        " ██████╔╝██║  ██║██║ ╚████║╚██████╗╚██████╔╝",
        " ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═════╝ "
    ]

    # ── Layout ─────────────────────────────────────────────────────

    @staticmethod
    def margem():
        return "  "

    @staticmethod
    def limpar():
        try:
            sistema = platform.system().lower()
            if "win" in sistema:
                os.system("cls")
            else:
                os.system("clear")
        except Exception:
            print("\033[H\033[2J", end="")

    @staticmethod
    def exibir_cabecalho(subtitulo=None):
        m = Banner.margem()
        print()

        meio = len(Banner.LOGO) // 2

        for i in range(len(Banner.LOGO)):
            sufixo = f"   {Banner.B_CINZA}{Banner.NOME}{Banner.RESET}" if i == meio else ""
            print(f"{m}{Banner.B_AZUL}{Banner.LOGO[i]}{Banner.RESET}{sufixo}")

        print(f"{m}{Banner.CINZA}{Banner.pad_esquerda(Banner.VERSAO, Banner.LARGURA)}{Banner.RESET}")
        print(f"{m}{Banner.CINZA}{Banner.SEPARADOR_DUPLO}{Banner.RESET}")

        if subtitulo and subtitulo.strip():
            print(f"{m}{Banner.B_AMAR}{Banner.centralizar(subtitulo, Banner.LARGURA)}{Banner.RESET}")

        print()

    @staticmethod
    def titulo_secao(texto):
        m = Banner.margem()
        li = Banner.LARGURA - 2

        print(f"{m}{Banner.CINZA}┌{'─' * li}┐{Banner.RESET}")
        print(
            f"{m}{Banner.CINZA}│{Banner.RESET}{Banner.B_CINZA}{Banner.centralizar(texto, li)}{Banner.RESET}{Banner.CINZA}│{Banner.RESET}")
        print(f"{m}{Banner.CINZA}└{'─' * li}┘{Banner.RESET}")
        print()

    @staticmethod
    def label_secao(texto):
        m = Banner.margem()
        print(f"{m}{Banner.B_AZUL}▸ {texto}{Banner.RESET}")
        print(f"{m}{Banner.CINZA}{Banner.SEPARADOR}{Banner.RESET}")

    @staticmethod
    def campo(label, valor, largura_label=16):
        m = Banner.margem()
        rotulo = f"{Banner.CINZA}{label:<{largura_label}}{Banner.RESET}"
        print(f"{m}  {rotulo}{valor}")

    # ── Feedback ────────────────────────────────────────────────────

    @staticmethod
    def sucesso(msg):
        print()
        print(f"{Banner.margem()}{Banner.B_VERDE}  ✔  {msg}{Banner.RESET}")

    @staticmethod
    def erro(msg):
        print()
        print(f"{Banner.margem()}{Banner.B_VERM}  ✖  {msg}{Banner.RESET}")

    @staticmethod
    def info(msg):
        print(f"{Banner.margem()}{Banner.B_CIANO}  ℹ  {msg}{Banner.RESET}")

    # ── Espaçamento ─────────────────────────────────────────────────

    @staticmethod
    def espaco():
        print()

    @staticmethod
    def linha_separadora():
        print(f"{Banner.margem()}{Banner.CINZA}{Banner.SEPARADOR}{Banner.RESET}")

    @staticmethod
    def aguarde_enter():
        print(f"{Banner.margem()}{Banner.CINZA}  Pressione ENTER para continuar...{Banner.RESET}", end="")
        input()

    # ── Formatação monetária ────────────────────────────────────────

    @staticmethod
    def formatar_dinheiro(valor):
        return f"R$ {valor:.2f}"

    @staticmethod
    def valor_positivo(valor):
        return f"{Banner.B_VERDE}{Banner.formatar_dinheiro(valor)}{Banner.RESET}"

    @staticmethod
    def valor_negativo(valor):
        return f"{Banner.B_VERM}{Banner.formatar_dinheiro(valor)}{Banner.RESET}"

    # ── Badges ──────────────────────────────────────────────────────

    @staticmethod
    def badge(texto):
        return f"{Banner.B_CIANO} {texto} {Banner.RESET}"

    @staticmethod
    def badge_verde(texto):
        return f"{Banner.B_VERDE} {texto} {Banner.RESET}"

    @staticmethod
    def badge_azul(texto):
        return f"{Banner.B_AZUL} {texto} {Banner.RESET}"

    # ── Helpers de cor ──────────────────────────────────────────────

    @staticmethod
    def ciano(s):
        return f"{Banner.B_CIANO}{s}{Banner.RESET}"

    @staticmethod
    def verde(s):
        return f"{Banner.B_VERDE}{s}{Banner.RESET}"

    @staticmethod
    def vermelho(s):
        return f"{Banner.B_VERM}{s}{Banner.RESET}"

    @staticmethod
    def amarelo(s):
        return f"{Banner.B_AMAR}{s}{Banner.RESET}"

    @staticmethod
    def cinza(s):
        return f"{Banner.CINZA}{s}{Banner.RESET}"

    @staticmethod
    def branco(s):
        return f"{Banner.B_CINZA}{s}{Banner.RESET}"

    @staticmethod
    def bold(s):
        return f"{Banner.BOLD}{s}{Banner.RESET}"

    # ── Helpers de layout ───────────────────────────────────────────

    @staticmethod
    def centralizar(texto, largura):
        if len(texto) >= largura:
            return texto
        pad = (largura - len(texto)) // 2
        return " " * pad + texto + " " * (largura - len(texto) - pad)

    @staticmethod
    def pad_esquerda(texto, largura):
        if len(texto) >= largura:
            return texto
        return " " * (largura - len(texto)) + texto

    # ── Prompt ──────────────────────────────────────────────────────

    @staticmethod
    def prompt(label):
        print(f"{Banner.margem()}  {Banner.B_AZUL}{label}{Banner.RESET} ", end="")