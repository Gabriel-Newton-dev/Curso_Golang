// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 20

package main

import "fmt"

// obs: para fazer esse programa temos que utilizar o for.
func main() {

	for i := 10; i <= 20; i++ {
		fmt.Println(i % 4)
	}

	fmt.Println("Agora iremos fazer um outro loop abaixo, porém incrementando de 3 em 3 até chegar a 99")

	// crie um outro loop que incremente de 3 em 3 até chegar 99

	for x := 0; x <= 99; x += 3 {
		fmt.Println(x)
	}

	// agora iremos fazer um loop que diminua de o valor 100 de 10 em 10 até chegar 0

	fmt.Println("agora iremos fazer um loop que diminua de o valor 100 de 10 em 10 até chegar 0")
	for y := 100; y >= 0; y -= 10 {
		fmt.Println(y)
	}

}
