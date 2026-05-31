from dataclasses import dataclass
from typing import List


@dataclass
class CadastroRequest:
    nome: str
    cpf: str
    senha: str
    tipoConta: int

@dataclass
class LoginRequest:
    cpf: str
    senha: str
    tipoConta: int

@dataclass
class PagamentoRequest:
    descricao: str
    valor: float

@dataclass
class RendimentoRequest:
    meses: int

@dataclass
class TransferenciaRequest:
    numDestino: int
    valor: float

@dataclass
class ValorRequest:
    valor: float

@dataclass
class CadastroResponse:
    mensagem: str
    numConta: int

@dataclass
class ContaResponse:
    numero: int
    titular: str
    saldo: float
    tipo: int
    limite: float
    taxa: float

@dataclass
class ExtratoResponse:
    numeroConta: int
    titular: str
    saldoAtual: float
    transacoes: List[TransacaoResponse]

@dataclass
class LoginResponse:
    mensagem: str
    nome: str
    numero: int

@dataclass
class MovimentacaoResponse:
    mensagem: str
    numeroConta: int
    tipoOperacao: int
    valorMovimentado: float
    saldoAtual: float

@dataclass
class PagamentoResponse:
    mensagem: str
    numeroConta: int
    descricao: str
    valorPago: float
    saldoAtual: float

@dataclass
class RendimentoResponse:
    numeroConta: int
    saldoAtual: float
    meses: int
    taxaMensal: float
    saldoProjetado: float
    rendimento: float

@dataclass
class TransacaoResponse:
    tipo: str
    valor: float
    descricao: str
    data: str

@dataclass
class TransferenciaResponse:
    mensagem: str
    contaOrigem: int
    contaDestino: int
    valorTransferido: float
    saldoAtualOrigem: float

@dataclass
class ContaLogada:
    numero: int
    nome: str
    cpf: str
    tipo: int
    saldo: float
    limite: float
    rendimento: float