package main

import "fmt"

func main() {

	canal := make(chan string, 200)
	canal <- "Olá mundo"
	canal <- "Programando em Golang"
	canal <- "Programando em Golang novamente"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
