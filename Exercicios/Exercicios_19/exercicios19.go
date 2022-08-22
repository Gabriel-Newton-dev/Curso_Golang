// - Crie uma slice usando make que possa conter todos os estados do Brasil.
//     - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás",
// "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba",
// "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul",
//  "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
// - Demonstre o len e cap da slice.
// - Demonstre todos os valores da slice sem utilizar range.

package main

import "fmt"

func main() {

	slice := make([]string, 26)

	slice = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(slice), cap(slice))

	// obs: nesse caso o exercicio pediu para não usar o Range.
	// for _, value := range slice {
	// 	fmt.Println(value)

	for i := 1; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
