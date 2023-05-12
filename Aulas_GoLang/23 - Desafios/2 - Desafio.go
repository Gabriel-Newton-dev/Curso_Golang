// Faça um programa em Golang que receba dois números e no final mostre a soma,
// subtração, multiplicação e a divisão dos números lidos.
// Os números deverão ser informados pelo usuário.

package main

import "fmt"

var number1, number2 float64

func main() {

	fmt.Print("Favor entre com o primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Favor entre com o segundo número: ")
	fmt.Scan(&number2)

	soma := number1 + number2
	substracao := number1 - number2
	multiplicacao := number1 * number2
	divisao := number1 / number2

	fmt.Printf("A soma é : %v \nA subtração é : %v \nA Multiplicação é : %v \nA divisão é : %v\n", soma, substracao, multiplicacao, divisao)
}
