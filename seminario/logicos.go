package main

import "fmt"

func main() {
	x, y := 16, 4
	conjuncao := x != y && x == 16
	disjuncao := x == 5 || y == 4
	negacao := !disjuncao

	fmt.Println("Conjunção: ", conjuncao)
	fmt.Println("Disjunção: ", disjuncao)
	fmt.Println("Negação: ", negacao)
}
