// - Crie dois slice contendo slices de strings ([][]string). Atribua valores a este slice
// multi-dimensional da seguinte maneira:
// - primeiro Slice : "Nome", "Sobrenome", "Hobby favorito"
// - Segundo Slice : de cachorro, raça e o que gosta de fazer
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

func main() {

	ss := [][]string{
		{"Joao", "Silva", "cantar igual gazela"},
		{"Maria", "Vai com as outras", "paquerar"},
		{"Roberta", "Jubiscreuda", "Rebolar na balada"},
	}

	s2 := [][]string{
		{"Rex", "Pitbull", "Morder"},
		{"Hulk", "American Bully", "Brincar"},
		{"Princesa", "Pastor Alemão", "caçar"},
	}

	for _, value := range ss {
		fmt.Println(value)
	}

	fmt.Println("")

	for _, valor := range s2 {
		fmt.Println(valor)
	}

}
