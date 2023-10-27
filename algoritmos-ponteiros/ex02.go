//Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro é par ou ímpar. A função deve
//atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.

package main

import "fmt"

func main() {
	var num int
	var ptr *int

	fmt.Println("O programa irá transformar os pares em 0 e ímpares em 1")
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	ptr = &num //a variavel ptr aponta para um endereco de memoria de num

	oddOrEven(ptr)
	fmt.Print("Novo valor do número: ", num)
}

func oddOrEven(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}
