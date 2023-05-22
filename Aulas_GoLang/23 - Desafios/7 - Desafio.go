// Crie um array de strings na linguagem Golang
// e use o laço for para acessar seus elementos de forma individual.
// Use função len(), para retornar a quantidade de itens no array.

package main

import (
	"fmt"
	"sort"
)

func main() {

	sliceString := []string{"Vectra", "Del Rey", "Civic", "Audi", "Omega", "Belina", "Ranger", "Gol", "Eco-Sport"}

	sort.Strings(sliceString)

	for key, value := range sliceString {
		fmt.Println(key, value)
	}

	fmt.Printf("Função Len: %d\n", len(sliceString))

}
