package main

import "fmt"

func main() {
	idade := 29

	if idade < 16 {
		fmt.Println("Não pode votar!")
	} else if idade < 18 || idade > 65 {
		fmt.Println("Voto opcional!")
	} else {
		fmt.Println("Voto obrigatório!")
	}
}
