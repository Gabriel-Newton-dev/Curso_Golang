// - Crie um programa que utilize a declaração switch, sem switch statement especificado.

package main

import "fmt"

func main() {

	flamengo := 5

	switch {
	case flamengo == 1:
		fmt.Println("Flamengo ganhou 1 jogo")
	case flamengo == 2:
		fmt.Println("Flamengo ganhou 2 jogos")
	case flamengo == 3:
		fmt.Println("Flamengo ganhou 3 jogos")
	case flamengo == 4:
		fmt.Println("Flamengo ganhou 4 jogos")
	case flamengo == 5:
		fmt.Println("Flamengo ganhou 5 jogos, ninguém segura o Mengão")
	}

}
