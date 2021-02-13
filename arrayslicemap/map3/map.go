package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 23234.88,
			"Guga Pereira":   13234.88,
		},
		"J": {
			"Julio Martins": 2134.99,
		},
		"V": {
			"Valentina Fragoso": 52342.99,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "J")

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println("Funcionarios que inicial com a letra", letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
