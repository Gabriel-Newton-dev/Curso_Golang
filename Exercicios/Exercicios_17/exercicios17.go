// - Crie um programa que utilize a declaração switch, onde o switch
// statement seja uma variável do tipo string com identificador "esporteFavorito".

package main

import "fmt"

func main() {

	esporteFavorito := "Futebol"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Meu esporte favorito é o Futebol, e irei jogar no Flamengo")
	case "Voley":
		fmt.Println("Meu esporte favorito é o Voley")
	case "Artes Marciais":
		fmt.Println("Meu esporte favorito é Artes Marciais")
	}

}
