package main

import (
	"alura-golang-oo/contas"
	"fmt"
)

func main() {
	contaAlex := contas.ContaCorrente{"Alex Moreno", 123, 12345678, 2000}
	contaAlex.Sacar(-1800)
	fmt.Println(contaAlex)
}
