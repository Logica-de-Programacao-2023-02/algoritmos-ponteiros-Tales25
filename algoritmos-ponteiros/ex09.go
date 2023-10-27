//Implemente uma função que receba um ponteiro para uma struct "Livro" com campos "Título", "Autor" e "Preço",
//e que adicione um desconto de 10% no preço do livro.

package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func main() {
	livro1 := Livro{
		Titulo: "Livro 1",
		Autor:  "Anonimo",
		Preco:  750,
	}

	fmt.Println(livro1.Preco)

	ptr := &livro1
	descontoLivro(ptr)
	fmt.Println(livro1.Preco)
}

func descontoLivro(ptr *Livro) {
	ptr.Preco = ptr.Preco * 90 / 100
}
