// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função,
// onde esta outra função faz uso de uma variável alem de seu scope.

package main

import "fmt"

func main() {

	a := i()
	b := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
