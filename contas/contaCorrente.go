package contas

import (
	"alura-golang-oo/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (cc *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) {
	if valorEhValido(valor) && valor < cc.saldo {
		cc.saldo -= valor
		contaDestino.Depositar(valor)
	}
}

func (cc *ContaCorrente) Sacar(valor float64) {
	if valorEhValido(valor) && valor < cc.saldo {
		cc.saldo -= valor
	}
}

func (cc *ContaCorrente) Depositar(valor float64) {
	if valorEhValido(valor) {
		cc.saldo += valor
	}
}

func valorEhValido(valor float64) bool {
	return valor > 0
}
