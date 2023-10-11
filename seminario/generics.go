package main

import (
	"fmt"
)

func menorValorInt(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func menorValorFloat(x float64, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(menorValorInt(7, 4))
	fmt.Println(menorValorFloat(0.3457, 0.3475))
}

// type tiposPermitidos interface {
// 	float32 | int
// }

// func menorValor[T tiposPermitidos](x T, y T) T {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// func main() {
// 	fmt.Println(menorValor[int](7, 4))
// 	fmt.Println(menorValor[float32](0.3457, 0.3475))
// }
