package main

import "fmt"

var number1, number2 float64
var operator string

func main() {

	fmt.Print("Enter with number first: ")
	fmt.Scan(&number1)

	fmt.Print("Enter with number second: ")
	fmt.Scan(&number2)

	fmt.Print("Enter with operator (+ | - | * | / ): ")
	fmt.Scan(&operator)

	if operator == "+" {
		fmt.Printf("%v + %v = %v\n", number1, number2, number1+number2)
	} else if operator == "-" {
		fmt.Printf("%v - %v = %v\n", number1, number2, number1-number2)
	} else if operator == "*" {
		fmt.Printf("%v * %v = %v\n", number1, number2, number1*number2)
	} else if operator == "/" {
		fmt.Printf("%v / %v = %v\n", number1, number2, number1/number2)
	} else {
		fmt.Println("Enter valid value :(")
	}
}
