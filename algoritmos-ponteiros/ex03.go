//Escreva uma função em Go que receba um ponteiro para uma string e atualize o valor da string para seu reverso.

package main

import "fmt"

func main() {
	str := "O rato roeu a roupa do Rei de Roma"
	ptr := &str

	fmt.Println(str)
	reverseString(str, ptr)
	fmt.Println(str)
}

func reverseString(str string, ptr *string) {
	reversedStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	*ptr = reversedStr
}
