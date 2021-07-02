package main

import "fmt"

type ContaCorrente struct {
	titula  string
	agencia int
	conta   int
	saldo   float64
}

func (cc *ContaCorrente) Sacar(valor float64) {
	if valor > 0 && valor < cc.saldo {
		cc.saldo -= valor
	}
}

func main() {
	contaAlex := ContaCorrente{"Alex Moreno", 123, 12345678, 2000}
	contaAlex.Sacar(-1800)
	fmt.Println(contaAlex)
}
