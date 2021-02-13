package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"Gustavo Fragoso":    20547.56,
		"Valentina Fragoso":  30459.45,
		"Gabriela Allegrini": 21457.55,
	}

	funcsESalarios["Rafael Filho"] = 664.99

	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
