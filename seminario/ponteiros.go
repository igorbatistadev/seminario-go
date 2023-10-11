package main

import "fmt"

func mudarNome(nome *string) {
	*nome = "Fulano"
}

func main() {
	nome := "Beltrano"
	var ponteiro *string = &nome
	// ponteiro := &nome

	fmt.Println("nome =", nome)
	fmt.Println("ponteiro =", ponteiro)

	fmt.Println("*ponteiro =", *ponteiro)

	mudarNome(&nome) // mudarNome(ponteiro)

	fmt.Print("\n\n")

	fmt.Println("nome =", nome)
	fmt.Println("ponteiro =", ponteiro)

	fmt.Println("*ponteiro =", *ponteiro)
}
