//Implemente uma função que receba um ponteiro para uma slice e seu
//tamanho e preencha-o com os n primeiros números primos.

package main

import "fmt"

func main() {
	sliceInt := []int{1, 2, 3, 4, 5, 4, 5, 5, 5, 5, 5, 6, 7, 8, 7, 8, 8}
	ptr := &sliceInt

	fmt.Println(sliceInt)

	//para fazer esse exercicio voce precisa de um booleano para verificar se o numero é primo ou nao, de uma variavel
	//(que comece em 2) para o numero e uma variavel de tamanho para preencher um slice com essa quantidade de numeros
	//primos, depois voce cria um loop de "enquanto o tamanho do seu slice for menor do que a variavel de tamanho",
	//voce parte do pressuposto que o numero é primo e atribui valor true a var booleana, entao inicia outro loop
	//de "i recebe 2 sempre que este loop for iniciado, enquanto i for menor do que o numero (inicialmente 2) adicione
	//um ao i (é resetado toda vez que entra no loop novamente)", cria um if onde "se o resto da divisao do numero
	//por i for igual a zero", a variavel booleana recebe false e voce aplica um break nesse segundo loop, depois se
	//o booleano for verdadeiro voce aplicar um append no seu slice (inicialmente vazio) e o numero, finalizando
	//voce adiciona um ao numero e enquanto o tamanho do seu slice for menor do que o tamanho desejado, ele ira
	//se repetir
	fillSliceWithPrimes(ptr)
	fmt.Println(sliceInt)
}

func fillSliceWithPrimes(ptr *[]int) {
	var isPrime bool
	var slicePrimes []int
	sliceInt := *ptr
	num := 2

	for len(slicePrimes) < len(sliceInt) {
		isPrime = true

		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			slicePrimes = append(slicePrimes, num)
		}

		num++
	}

	*ptr = slicePrimes
}
