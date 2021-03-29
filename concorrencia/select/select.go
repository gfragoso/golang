package main

import (
	"fmt"
	"time"

	"github.com/gfragoso/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle especifica para concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Second):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.ifood.com.br",
		"https://www.google.com",
		"https://www.youtube.com",
	)
	fmt.Println(campeao)
}
