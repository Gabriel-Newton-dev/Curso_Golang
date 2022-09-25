package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)

	anonima := func(palavra string) string {
		return fmt.Sprintf("No caso em tela fizemos uma : %s", palavra)
	}("Função anônima")

	fmt.Println(anonima)

}
