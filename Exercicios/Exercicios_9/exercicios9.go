// - Crie um programa que:
//     - Atribua um valor int a uma variável
//     - Demonstre este valor em decimal, binário e hexadecimal
//     - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
//     - Demonstre esta outra variável em decimal, binário e hexadecimal

package main

import "fmt"

var x int = 200

func main() {

	fmt.Printf("Valor em Decimal : %d\nValor em Binário %b\nValor em Hexadecimal %#x\n", x, x, x)

	// Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável

	y := x << 1
	fmt.Println("Valor abaixo refere-se a Y")
	fmt.Printf("Valor em Decimal : %d\nValor em Binário %b\nValor em Hexadecimal %#x", y, y, y)

}
