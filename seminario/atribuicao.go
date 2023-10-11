package main

import "fmt"

func main() {
	x, y := 16, 4
	x = 20
	y += 10
	y -= 10
	y *= 10
	y /= 10
	x %= 2

	fmt.Println("X:", x)
	fmt.Println("Y:", y)
}
