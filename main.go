package main

import "fmt"

type ContaCorrente struct {
	titula  string
	agencia int
	conta   int
	saldo   float64
}

func main() {
	contaAlex := ContaCorrente{"Alex Moreno", 123, 12345678, 2000}
	fmt.Println(contaAlex)
}
