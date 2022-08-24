// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

package main

import "fmt"

func main() {

	x := map[string][]string{
		"Seu madruga do Chaves": {"Bater no chaves", "Não pagar o aluguel", "Apanhar da dona florinda"},
		"Chaves da Silva":       {"Comer sanduiche de presunto", "Jogar bola", "Não tomar banho"},
		"Dona Florinda":         {"Bater no seu madruga", "Falar mal da gentalha", "Namorar o professor Girafales"},
	}

	for key, value := range x {
		fmt.Println(key)
		for k, v := range value {
			fmt.Println("\t", k, " - ", v)
		}
	}
}
