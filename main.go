package main

import (
	"fmt"
	"go-bank/contas"
)



func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGabriel := contas.ContaCorrente{Titular: "Gabriel", Saldo: 100}

	staus := contaDaSilvia.Transferir(-200, &contaDoGabriel)

	fmt.Println(staus)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGabriel)
}