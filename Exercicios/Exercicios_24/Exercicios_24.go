// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice
// que contem os sabores de sorvete.

package main

import "fmt"

type pessoa struct {
	Nome            string
	SobreNome       string
	SaboresSorvetes []string
}

func main() {

	Joao := pessoa{
		Nome:            "Joao",
		SobreNome:       "Pedro",
		SaboresSorvetes: []string{"Flocos", "Morango", "Chocolate"},
	}

	Maria := pessoa{
		Nome:            "Maria",
		SobreNome:       "Betania",
		SaboresSorvetes: []string{"Passas ao rum", "Abacaxi", "Napolitano"},
	}

	fmt.Println(Joao.Nome, Joao.SobreNome)
	for _, valor := range Joao.SaboresSorvetes {
		fmt.Println("\t", valor)
	}
	fmt.Println(Maria.Nome, Maria.SobreNome)
	for _, value := range Maria.SaboresSorvetes {
		fmt.Println("\t", value)
	}

}
