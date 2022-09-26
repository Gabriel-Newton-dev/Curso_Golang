// a função pré-definida init() faz com que uma parte do código execute antes de qualquer outra parte do seu pacote.

package main

import "fmt"

var n int

func main() {
	fmt.Println("Package main")
	fmt.Println(n)
}

func init() {
	fmt.Println("Init faz com que uma parte do código execute antes de qq outra parte do seu pacote") // mesmo estando aqui depois da main ele irá executar primeiro que todos.
	n = 10
}
