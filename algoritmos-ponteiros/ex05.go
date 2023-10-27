//Escreva uma função em Go que receba um ponteiro para uma variável float64 e atualize o valor da variável para a média
//aritmética entre o seu valor atual e o valor da constante Pi.

package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	var ptr *float64

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	ptr = &num

	calculateMediaWithPi(ptr)
	fmt.Print("A média entre o número e PI é: ", num)
}

func calculateMediaWithPi(ptr *float64) {
	*ptr = (*ptr + math.Pi) / 2
}
