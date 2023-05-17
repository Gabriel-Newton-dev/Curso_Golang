// Escrever um programa em Golang para ler uma temperatura em graus Celsius e
// apresentá-la convertida em graus Fahrenheit. A fórmula de conversão é:
// F = ((C * 9) / 5) + 32, sendo F a temperatura em Fahrenheit e C a temperatura
// em Celsius.

// Sua saída deverá ser parecida com:

// Informe a temperatura em Celsius: 40
// 40.0 graus Celcius é igual à 104.0 graus Fahrenheit.

package main

import "fmt"

var grauCelcius float64

func main() {

	fmt.Print("Entre com a temperatura em Graus Celcius: ")
	fmt.Scan(&grauCelcius)

	f := ((grauCelcius * 9) / 5) + 32
	fmt.Printf("%v graus Celcius é igual à %v graus Fahrenheit.\n", grauCelcius, f)

}
