package main

import "fmt"

func main() {
	x, y := 16, 4
	igualdade := x == y
	diferenca := x != y
	maior := x > y
	menor := x < y
	maior_igual := x >= y
	menor_igual := x <= y

	fmt.Println("Igualdade: ", igualdade)
	fmt.Println("Diferença: ", diferenca)
	fmt.Println("Maior: ", maior)
	fmt.Println("Menor: ", menor)
	fmt.Println("Maior ou igual: ", maior_igual)
	fmt.Println("Menor ou igual: ", menor_igual)
}
