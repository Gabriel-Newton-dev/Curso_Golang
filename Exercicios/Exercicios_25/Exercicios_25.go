// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

package main

import "fmt"

type pessoa struct {
	Nome            string
	SobreNome       string
	SaboresSorvetes []string
}

func main() {

	// jeito mais fácil de usar um make no map

	meuMapa := make(map[string]pessoa)

	meuMapa["1"] = pessoa{
		Nome:            "Joao",
		SobreNome:       "Pedro",
		SaboresSorvetes: []string{"Flocos", "Morango", "Chocolate"},
	}

	meuMapa["2"] = pessoa{
		Nome:            "Maria",
		SobreNome:       "Betania",
		SaboresSorvetes: []string{"Passas ao rum", "Abacaxi", "Napolitano"},
	}

	for _, value := range meuMapa {
		fmt.Println(value)
	}

}
