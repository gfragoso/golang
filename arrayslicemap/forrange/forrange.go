package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta

	for i, numumero := range numeros {
		fmt.Printf("%d) %d\n", i, numumero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
