package contas

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (cc *ContaCorrente) Sacar(valor float64) {
	if valor > 0 && valor < cc.saldo {
		cc.saldo -= valor
	}
}
