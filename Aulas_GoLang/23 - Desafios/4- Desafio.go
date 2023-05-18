// Escreva um programa Golang que leia duas notas (como float64), calcule e mostre a
// média aritmética e uma mensagem de acordo com as seguintes regras:

//1) Se a média for inferior a 4,0 escreva "Reprovado";
// 2) Se a média for igual ou superior a 4,0 e inferior a 7,0 escreva "Exame";
// 3) Se a média for igual ou superior a 7,0 escreva "Aprovado".

package main

import "fmt"

var nota1 float64
var nota2 float64

func main() {

	fmt.Print("Entre com a primeira nota: ")
	fmt.Scan(&nota1)

	fmt.Print("Entre com a segunda nota: ")
	fmt.Scan(&nota2)

	resultado := (nota1 + nota2) / 2

	if resultado < 4.0 {
		fmt.Printf("%v + %v = Média %v - Reprovado\n", nota1, nota2, resultado)
	} else if resultado >= 4.0 && resultado < 7.0 {
		fmt.Printf("%v + %v = Média %v - Exame\n", nota1, nota2, resultado)
	} else if resultado >= 7.0 {
		fmt.Printf("%v + %v = Média %v - Aprovado\n", nota1, nota2, resultado)
	}

}
