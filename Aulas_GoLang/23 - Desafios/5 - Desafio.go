// Escreva um programa Go que leia o peso de uma pessoa na Terra e retorne
// o seu peso na Lua. Lembre-se da seguinte fórmula:
// Aqui nós estamos definindo a força da gravidade na Terra como 9,81 m/s2 e
// a força da gravidade na Lua como 1,622 m/s2.

package main

import "fmt"

var pesoPessoa float64

func main() {

	fmt.Print("Quantos Kilos vc pesa ? ")
	fmt.Scan(&pesoPessoa)

	pesoNaLua := pesoPessoa / 9.81 * 1.622
	fmt.Printf("Seu peso na lua é de %.2f KG.\n", pesoNaLua)
}
