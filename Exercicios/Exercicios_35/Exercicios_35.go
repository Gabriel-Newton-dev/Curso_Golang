// - Callback: passe uma função como argumento a outra função.

package main

import "fmt"

func main() {

	funcao2(funcao1)

}

func funcao1() {
	fmt.Println("Acabamos de fazer um callback em Golang!")
}

func funcao2(x func()) {
	fmt.Printf("Preste atenção: ")
	x()
}
