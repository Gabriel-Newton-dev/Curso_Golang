package main

import (
	"fmt"
)

func main() {

	go escrever("Estamos usando Goroutines")
	escrever("Dessa forma que aplicamos o Goroutines")

}

func escrever(texto string) {
	fmt.Println(texto)
	// time.Sleep(time.Second)
}
