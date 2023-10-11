package main

import (
	"fmt"
	"strings"
)

func boasvindas() {
	fmt.Println("*****************************************")
	fmt.Println("*** BEM VINDO AO CADASTRO DE CLIENTES ***")
	fmt.Println("*****************************************")
}

func menu() {
	fmt.Println("\nMenu: ")
	fmt.Println("1 - Cadastrar cliente. ")
	fmt.Println("2 - Listar nome e idade de todos clientes. ")
	fmt.Println("3 - Buscar cliente por nome. ")
	fmt.Println("4 - Sair do sistema. ")
}

func solicitaOpcao() int {
	var opcao int
	fmt.Print("Digite sua opção: ")
	fmt.Scanln(&opcao)
	return opcao
}

type Cliente struct {
	Nome  string
	Idade int
	Cpf   string
}

func (c Cliente) toString() string {
	return fmt.Sprintf("Cliente: %s, Idade: %d", c.Nome, c.Idade)
}

func cadastrarCliente(lista *[]Cliente) {
	var cliente *Cliente
	cliente = new(Cliente)
	fmt.Println("Qual o nome do cliente?")
	fmt.Scanln(&cliente.Nome)
	fmt.Println("Qual a idade?")
	fmt.Scanln(&cliente.Idade)
	fmt.Println("Qual o CPF?")
	fmt.Scanln(&cliente.Cpf)

	fmt.Println("Dados =", *cliente)
	*lista = append(*lista, *cliente)
	fmt.Println("Cliente adicionado!")

	menu()
}

func listarTodosClientes(lista []Cliente) {
	if len(lista) == 0 {
		fmt.Println("Nenhum cliente cadastrado!")
		return
	}

	for _, cliente := range lista {
		fmt.Println(cliente.toString())
	}
}

func buscarPorNome(lista []Cliente) {
	nomeDesejado := ""
	fmt.Println("Qual o nome que deseja buscar?")
	fmt.Scanln(&nomeDesejado)

	listaClientesEncontrados := []Cliente{}

	for _, cliente := range lista {
		if strings.Contains(strings.ToLower(cliente.Nome), strings.ToLower(nomeDesejado)) {
			listaClientesEncontrados = append(listaClientesEncontrados, cliente)
		}
	}

	if len(listaClientesEncontrados) > 0 {
		fmt.Println("Clientes encontrados:")
		for _, cliente := range listaClientesEncontrados {
			fmt.Println(cliente.toString())
		}
	} else {
		fmt.Println("Nenhum cliente encontrado!")
	}

	menu()
}

func main() {

	boasvindas()

	var listaClientes []Cliente

	for {
		menu()
		opcao := solicitaOpcao()

		switch opcao {
		case 1:
			cadastrarCliente(&listaClientes)
		case 2:
			listarTodosClientes(listaClientes)
		case 3:
			buscarPorNome(listaClientes)
		case 4:
			fmt.Println("Saindo do sistema... ")
			return
		default:
			fmt.Println("Opção inválida, vamos tentar novamente")
		}
	}

}
