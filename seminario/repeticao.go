package main

import "fmt"

func main() {
	contador := 1
	for {
		fmt.Println(contador)
		contador++
		if contador == 5 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
