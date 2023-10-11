package main

import "fmt"

func main() {
	x, y := 16, 4
	soma := x + y
	subtracao := x - y
	multiplicacao := x * y
	divisao := x / y
	resto := x % y
	fmt.Println("Soma: ", x, "+", y, "=", soma)
	fmt.Println("Subtração: ", x, "+", y, "=", subtracao)
	fmt.Println("Multiplicação: ", x, "+", y, "=", multiplicacao)
	fmt.Println("Divisão: ", x, "+", y, "=", divisao)
	fmt.Println("Resto: ", x, "+", y, "=", resto)
	x++
	y--
	fmt.Println("Incremento de X: ", x)
	fmt.Println("Decremento de Y: ", y)
}
