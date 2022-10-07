// Faça uma calculadora usando Golang, usando as 4 operações da matemática.

package main

import "fmt"

var number1 int
var number2 int
var operadores string

func main() {

	fmt.Printf("Digite o primeiro número por favor: ")
	fmt.Scan(&number1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&number2)

	fmt.Print("Qual operação deseja fazer (+ | - | *  | / | %) :")
	fmt.Scan(&operadores)

	switch operadores {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		fmt.Println(number1 / number2)
	case "%":
		fmt.Println(number1 % number2)
	}

}
