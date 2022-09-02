// faça a calculadora, com os seguintes operadores ( +, -, * e / )
// possibilitando o usuário de interagir com a mesma.

package main

import "fmt"

var number1 float64
var number2 float64
var operadores string

func main() {

	fmt.Print("Entre com o primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Entre com o segundo número: ")
	fmt.Scan(&number2)

	fmt.Print("Qual operação aritmética deseja fazer ( +, -, *, / ) : ")
	fmt.Scan(&operadores)

	switch operadores {
	case "+":
		fmt.Printf("O resultado da soma é: %v \n", number1+number2)
	case "-":
		fmt.Printf("O resultado da subtração é: %v \n", number1-number2)
	case "*":
		fmt.Printf("O resultado da Multiplicação é: %v \n", number1*number2)
	case "/":
		fmt.Printf("O resultado da divisão é: %v ", number1/number2)
		// case "%":
		// 	fmt.Println(number1 % number2)
	}
}
