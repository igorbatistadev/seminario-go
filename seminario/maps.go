package main

import "fmt"

func main() {
	aluno_nota := map[string]float32{
		"José":  9.2,
		"Pedro": 8.7,
		"Maria": 9.8,
	}
	fmt.Println(aluno_nota)
	fmt.Println("Nota José: ", aluno_nota["José"])

	// Adicionando
	aluno_nota["Alfredo"] = 7
	fmt.Println(aluno_nota)

	// Deletando
	delete(aluno_nota, "Alfredo")
	fmt.Println(aluno_nota)

	// Verificando
	nota, is_existe := aluno_nota["Alfredo"]
	if is_existe {
		fmt.Println("Nota:", nota)
	} else {
		fmt.Println("Alfredo não existe!")
	}

	nota, is_existe = aluno_nota["Maria"]
	if is_existe {
		fmt.Println("Nota:", nota)
	} else {
		fmt.Println("Maria não existe!")
	}
}
