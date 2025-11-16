package main

import (
	"fmt"
	"go-bank/clientes"
	"go-bank/contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.123.123-12", "Desenvolvedor Go"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	fmt.Println(contaDoBruno)
}

