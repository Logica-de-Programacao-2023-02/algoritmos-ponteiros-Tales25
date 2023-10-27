//Escreva uma função em Go que receba um ponteiro para um struct Conta com campos saldo e titular,
//e adicione um valor ao saldo da conta.

package main

import "fmt"

type Conta struct {
	Titular string
	Saldo   float64
}

func main() {
	conta1 := Conta{
		Titular: "Tales",
	}

	fmt.Println(conta1)

	ptr := &conta1
	addSaldo(ptr)

	fmt.Print(conta1)
}

func addSaldo(ptr *Conta) {
	fmt.Print("Quanto você tem na conta? ")
	fmt.Scan(&ptr.Saldo)
}
