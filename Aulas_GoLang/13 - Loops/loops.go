// Para lembrar toda vez que falamos de Loop temos que lembrar do FOR.

package main

import (
	"fmt"
	"time"
)

func main() {

	// iremos fazer um loop, incrementando até chegar o valor o qual desejamos, que será 20

	i := 0
	for i < 3 {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	// podemos fazer dessa forma tb, sem declarar o valor anteriormente, lembrando sempre de usar o FOR que é o loop.

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// quando é com strings podemos fazer dessa forma,

	nomes := []string{"Gabriel", "Newton", "O mestre", "de Golang"}

	for _, valor := range nomes {
		fmt.Println(valor)
	}

	// outra forma de fazer o loop, seria usando Range, para imprimir as palavras com key(que iremos colocar indice)
	// e value ( que iremos colocar com nome letra)

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) // para imprimir as letras temos que colocar string na frente da value(letra)
	}

	mapa := map[string]string{
		"Nome":      "Gabriel",
		"Sobrenome": "Newton",
	}
	for key, value := range mapa {
		fmt.Println(key, value)
	}

	// Dessa forma iremos fazer um for e executar infinitamente.
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}
