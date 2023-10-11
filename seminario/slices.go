package main

import (
	"fmt"
)

func main() {
	var notas []float32
	notas = append(notas, 6.5)
	notas = append(notas, 7.5)
	fmt.Println(notas)

	nomes := []string{"Fulano", "Beltrano"}
	fmt.Println(nomes)

	idades := make([]int, 5)
	idades[0] = 21
	idades[1] = 19
	idades = append(idades, 33)
	fmt.Println(idades)
}
