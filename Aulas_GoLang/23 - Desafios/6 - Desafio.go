// Escreva um programa Golang que leia a diagonal maior e
// a diagonal menor e calcule a área do losango. Sua saída deverá ser parecida com:

// Informe a medida da diagonal maior: 5
// Informe a medida da diagonal menor: 10
// A área (em metros quadrados) do losango é: 25

package main

import (
	"fmt"
)

var medida1, medida2 int

func main() {

	fmt.Printf("Informe a medida da diagonal maior: ")
	fmt.Scan(&medida1)

	fmt.Printf("Informe a medida da diagonal menor: ")
	fmt.Scan(&medida2)

	areaDoLosango := medida1 * medida2 / 2
	fmt.Printf("A área desse losango é de : %d \n", areaDoLosango)
}
