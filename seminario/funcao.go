package main

import "fmt"

func intro() {
	fmt.Println("Olá, vamos realizar uma soma!")
}

func soma(x int, y int) int {
	return x + y
}

func main() {
	intro()
	soma := soma(5, 3)
	fmt.Println("O resultado da soma é", soma)
}
