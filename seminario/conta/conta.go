package conta

import (
	"fmt"
	"seminario/titular"
)

type Conta struct {
	Titular       titular.Titular
	NumeroConta   int
	NumeroAgencia int
	Saldo         float64
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
