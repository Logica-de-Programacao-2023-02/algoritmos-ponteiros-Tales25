//Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n. A função deve atualizar o valor do
//inteiro para a soma dos n primeiros números naturais.

package main

import "fmt"

func main() {
	var n int
	var ptr *int

	fmt.Print("Digite um número (!=0): ")
	fmt.Scan(&n)
	ptr = &n

	fmt.Printf("A soma dos %d(n) primeiros números naturais", n)
	somaNatuaris(n, ptr)
	fmt.Printf(" é igual a %d", n)
}

func somaNatuaris(n int, ptr *int) {
	//zerei o valor inicial do ponteiro pq comeca com o valor de n
	*ptr = 0

	for i := 0; i < n; i++ {
		*ptr += i
	}
}
