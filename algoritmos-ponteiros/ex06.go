//Escreva uma função em Go que receba um ponteiro para um struct Livro com campos título e autor, e altere o título do
//livro para "Desconhecido" se o autor for "Anônimo".

package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func main() {
	Book1 := Book{
		Title:  "Work 4h/Week",
		Author: "Timothy Ferriss",
	}

	Book2 := Book{
		Title:  "48 Laws of Power",
		Author: "Anonymous",
	}

	fmt.Println(Book1)
	fmt.Println(Book2)

	ptr1, ptr2 := &Book1, &Book2
	correctBookInfo(ptr1)
	correctBookInfo(ptr2)

	fmt.Println(Book1)
	fmt.Println(Book2)
}

func correctBookInfo(ptr *Book) {
	//por algum motivo nao se usa o * para alterar o valor (ex.: *ptr.Author)
	if ptr.Author == "Anonymous" {
		ptr.Title = "Unknown"
	}
}
