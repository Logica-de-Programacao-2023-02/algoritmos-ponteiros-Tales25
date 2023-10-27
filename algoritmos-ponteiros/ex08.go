//Crie uma função que receba um ponteiro para uma variável como argumento e modifique o valor
//da variável dentro da função. Certifique-se de que o ponteiro aponte para uma
//área de memória válida e que a memória seja liberada após o uso.

//nao sei se oq era pra fazer está certo nem tenho ideia do que seja liberar memoria

package main

import "fmt"

func main() {
	var num int
	var ptr *int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	ptr = &num

	dobrarValor(ptr)

	fmt.Print(num)
}

func dobrarValor(ptr *int) {
	if ptr == nil {
		fmt.Println("Pointer is nil")
	} else {
		*ptr *= 2
	}
}
