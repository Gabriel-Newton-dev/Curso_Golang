// Crie uma nota numérica para condicional de nota por letra (ou vice-versa).

package main

import "fmt"

func main() {

	nota := 10.0

	A := nota

	fmt.Println()

	if nota == A {
		fmt.Println("Parabéns sua nota foi 10.0, vc está média")
	} else {
		fmt.Println("Conceito B em diante.")
	}

}
