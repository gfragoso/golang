package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func (p pedido) qtdeItens() int {
	total := 0
	for _, item := range p.itens {
		total += item.qtde
	}
	return total
}

func main() {

	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtde: 2, preco: 12.10},
			{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	fmt.Println(pedido)
	fmt.Printf("Quantidade total do itens é %v\n", pedido.qtdeItens())
	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
}
