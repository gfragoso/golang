package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// map deve ser inicializado

	aprovados := make(map[int]string)

	aprovados[32333333333] = "Maria"
	aprovados[13123123123] = "Gustavo"
	aprovados[12312312323] = "Valentina"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[13123123123])
	delete(aprovados, 32333333333)

	fmt.Println(aprovados)
}
