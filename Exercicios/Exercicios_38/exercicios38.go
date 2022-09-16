// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

package main

import (
	"fmt"
)

func main() {

	si := []int{1, 2, 3, 4, 5, 6, 82}
	si2 := []int{1, 22, 3, 4, 76, 82}

	fmt.Println(funcao1(si...))
	fmt.Println(funcao2(si2))
}

func funcao1(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func funcao2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
