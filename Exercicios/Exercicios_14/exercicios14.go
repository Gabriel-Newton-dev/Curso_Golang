// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import "fmt"

// obs: para fazer esse programa temos que utilizar o for.
func main() {

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}

}
