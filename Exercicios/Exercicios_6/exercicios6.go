// - Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

package main

import "fmt"

func main() {

	x := 100

	fmt.Printf("%d\n", x) // Decimal
	fmt.Printf("%b\n", x) //binario
	fmt.Printf("%#x", x)  //hexadecimal - por convenção usamos #.

}
