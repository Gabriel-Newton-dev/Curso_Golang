// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os
// elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(funcao3(slice...))

	slice2 := []int{60, 70, 80, 90, 100}
	fmt.Println(funcao4(slice2))

}

func funcao3(x ...int) int {
	total := 0
	for _, value := range x {
		total += value // mesma coisa que total = total + value
	}
	return total
}

func funcao4(y []int) int {
	total := 0
	for _, valor := range y {
		total += valor
	}
	return total
}
