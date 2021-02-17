package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	var coisa interface{}
	fmt.Println(coisa)

	var otacoisa any = "ASD"
	fmt.Println(otacoisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Expplorando a linguagem do Google"}
	fmt.Println(coisa2)
}
