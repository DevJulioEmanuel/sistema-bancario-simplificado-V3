from models import ContaLogada

usuario_atual: ContaLogada = None

def iniciar(conta: ContaLogada):
    global usuario_atual
    usuario_atual = conta

def encerrar():
    global usuario_atual
    usuario_atual = None

def esta_logado() -> bool:
    return usuario_atual is not None