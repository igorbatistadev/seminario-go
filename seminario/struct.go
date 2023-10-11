package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero string
	Bairro string
	Cidade string
	Estado string
	Cep    string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Cpf      string
	Endereco Endereco
}

func main() {
	joao := Pessoa{Nome: "João", Idade: 22, Cpf: "12345678900"}
	endereco_joao := Endereco{"Rua Tapajós", "28B", "Interlagos", "Linhares", "ES", "29903-700"}
	joao.Endereco = endereco_joao
	fmt.Println(joao)

	var maria *Pessoa
	maria = new(Pessoa)
	maria.Nome = "Maria"
	fmt.Println(maria)
}
