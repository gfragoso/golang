package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaPraConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaPraConceito(9.8))
	fmt.Println(notaPraConceito(7.9))
	fmt.Println(notaPraConceito(6.9))
	fmt.Println(notaPraConceito(4.8))
	fmt.Println(notaPraConceito(2.8))
}
