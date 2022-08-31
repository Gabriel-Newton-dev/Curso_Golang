// - Utilize a declaração defer de maneira que demonstre que sua execução
// só ocorre ao final do contexto ao qual ela pertence.

package main

import "fmt"

type Notebooks struct {
	Marca       string
	Modelo      string
	Processador string
	Valor       float64
}

func main() {

	Dell := Notebooks{
		Marca:       "Dell",
		Modelo:      "Xps",
		Processador: "I9",
		Valor:       9560.99,
	}

	Macbook := Notebooks{
		Marca:       "Apple",
		Modelo:      "Macbook Pro",
		Processador: "M1",
		Valor:       12.650,
	}

	defer fmt.Println(Dell) // com defer na frente será impresso por último.
	fmt.Println(Macbook)

}
