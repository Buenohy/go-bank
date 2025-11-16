package main

import (
	"fmt"
	"go-bank/contas"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(55)
	fmt.Println(contaDoDenis.ObterSaldo())
}

