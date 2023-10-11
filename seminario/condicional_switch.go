package main

import "fmt"

func main() {
	opcao := 2

	switch opcao {
	case 1:
		fmt.Println("Vamos cadastrar")
	case 2:
		fmt.Println("Vamos editar")
	case 3:
		fmt.Println("Vamos visualizar")
	case 4:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Opção inválida")
	}
}
