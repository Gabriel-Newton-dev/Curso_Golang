// - Usando uma literal composta:
//      - Crie um array que suporte 5 valores to tipo int
//      - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.

package main

import "fmt"

func main() {

	array := [5]int{10, 20, 30, 40, 50}

	for key, value := range array {
		fmt.Println(key, value)
	}
	fmt.Printf("%T", array) // demonstrar o tipo do array

}
