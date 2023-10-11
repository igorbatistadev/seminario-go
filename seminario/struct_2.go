package main

import (
	"fmt"
)

type Conta struct {
	Titular       Titular
	NumeroConta   int
	NumeroAgencia int
	Saldo         float64
}

type Titular struct {
	Nome string
	Cpf  string
}

func (c *Conta) Sacar(valorSaque float64) {
	if valorSaque <= c.Saldo {
		c.Saldo -= valorSaque
		fmt.Println("Saque efetuado com sucesso!")
	} else {
		fmt.Println("Saldo Insuficiente!")
	}
}

func (c *Conta) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		fmt.Println("Deposito realizado com sucesso!")
	} else {
		fmt.Println("Valor do deposito não pode ser negativo!")
	}
}

func main() {
	joao := Titular{Nome: "João", Cpf: "12345678900"}
	contaJoao := Conta{Titular: joao, NumeroConta: 1234567, NumeroAgencia: 0772, Saldo: 5000}
	fmt.Println(contaJoao)
	contaJoao.Sacar(1000)
	fmt.Println(contaJoao)
}
