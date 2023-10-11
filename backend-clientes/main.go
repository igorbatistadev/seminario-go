package main

import (
	"backend/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando servidor!")
	routes.HandleRequest()
}
