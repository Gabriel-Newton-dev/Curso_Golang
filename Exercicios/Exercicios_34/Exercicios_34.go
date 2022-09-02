// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.

package main

import "fmt"

func main() {

	x := retornaUmaFuncao()
	x()
}

func retornaUmaFuncao() func() {
	return func() {
		fmt.Println("Essa é uma função que retorna outra funçã0.")
	}

}
