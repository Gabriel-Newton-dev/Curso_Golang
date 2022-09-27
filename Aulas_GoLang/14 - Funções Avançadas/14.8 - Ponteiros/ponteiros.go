package main

import "fmt"

func inverterSinal(x int) int {
	return x * -1 // aqui coloquei ponteiro para inverter o sinal do int
}

func inverterSinalComPonteiro(y *int) { // podemos fazer dessa forma tb.
	*y = *y * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
