package main

import "fmt"

func notaPraConceiro(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaPraConceiro(9.8))
	fmt.Println(notaPraConceiro(6.9))
	fmt.Println(notaPraConceiro(2.1))
}
