package main

import (
	"alura-golang-oo/clientes"
	"alura-golang-oo/contas"
	"fmt"
)

func main() {
	alex := clientes.Titular{"Alex Moreno", "12312312312"}
	contaAlex := contas.ContaCorrente{Titular: alex, Agencia: 123, Conta: 12345678}
	contaAlex.Depositar(3000)
	contaAlex.Sacar(200)

	talita := clientes.Titular{"Talita Costa", "78978978965"}
	contaTalta := new(contas.ContaCorrente)
	contaTalta.Agencia = 332
	contaTalta.Conta = 14567
	contaTalta.Titular = talita

	contaAlex.Transferir(800, contaTalta)

	fmt.Println(contaAlex)
	fmt.Println(contaTalta)

}
