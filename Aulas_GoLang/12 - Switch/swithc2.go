// Fazer uma calculadora usando Switch

package main

import "fmt"

var number1 float64
var number2 float64
var operator string

func main() {

	fmt.Print("Please enter the first number: ")
	fmt.Scan(&number1)

	fmt.Print("Please enter the second number: ")
	fmt.Scan(&number2)

	fmt.Print("Enter the operator (+ | - | * | / ): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("%v + %v = %v\n", number1, number2, number1+number2)
	case "-":
		fmt.Printf("%v - %v = %v\n", number1, number2, number1-number2)
	case "*":
		fmt.Printf("%v * %v = %v\n", number1, number2, number1*number2)
	case "%":
		fmt.Printf("%v / %v = %v\n", number1, number2, number1/number2)
	}
}
