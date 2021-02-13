package main

import "fmt"

func notaPraConceiro(nota float64) string {

	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaPraConceiro(9.8))
	fmt.Println(notaPraConceiro(6.9))
	fmt.Println(notaPraConceiro(2.1))
}
