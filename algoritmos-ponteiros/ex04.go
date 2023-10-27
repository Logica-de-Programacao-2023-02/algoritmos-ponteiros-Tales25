//Escreva uma função em Go que receba um ponteiro para uma variável inteira e atualize
//o valor da variável para a soma dos valores dos seus dois últimos dígitos
//(por exemplo, se o valor inicial da variável for 1234, o novo valor será 3+4=7).

package main

import (
	"fmt"
)

func main() {
	var num int
	var ptr *int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	ptr = &num

	sumLastsNumbers(ptr)
	fmt.Printf("Agora o número vale %d (soma dos 2 últimos dígitos)", num)
}

func sumLastsNumbers(ptr *int) {
	//MINHA TENTATIVA
	//strInt := strconv.Itoa(*ptr)
	//sliceStrInt := strings.Split(strInt, "")
	//
	//index := len(sliceStrInt) - 1
	//
	//penultimateStr := sliceStrInt[index-1]
	//lastStr := sliceStrInt[index]
	//
	//penultimateNum, _ := strconv.Atoi(penultimateStr)
	//lastNum, _ := strconv.Atoi(lastStr)
	//
	//*ptr = penultimateNum + lastNum

	//JEITO DO THIAGO
	//Ex.: 1234 (34-4)/10 + 4 => 30/10 + 4 => 3 + 4 = 7
	*ptr = (*ptr%100-*ptr%10)/10 + *ptr%10
}
