// - Crie um loop utilizando a sintaxe: for condition {}
// - Utilize-o para demonstrar os anos desde que você nasceu.

package main

import "fmt"

func main() {

	anoQueNasci := 1985
	anoQueQueroCalcular := 2022

	for loop := anoQueNasci; loop <= anoQueQueroCalcular; loop++ {
		fmt.Println(loop)
	}
}

// Poderiamos fazer de uma forma mais simples usando uma variavel i

// for i := 1985; i <= 2022; i ++ ( ++ incrementa até chegar a 2022){
// fmt.Println(i)}
