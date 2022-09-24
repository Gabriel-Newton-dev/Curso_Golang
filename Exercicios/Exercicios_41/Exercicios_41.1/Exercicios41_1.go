// Crie um struct com nome carros
// crie uma funcao chamada mudeMe usando um *carros que altere o Marcae o valor do carro, usando um ponteiro
// podemos usar tanto o atalho c1.Marca quanto o tradicional (*c1).Marca
// - Crie um valor do tipo carros;
// - Use a função mudeMe e demonstre o resultado.
// demostre o valor em json

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Carros struct {
	Nome  string  `json:"Nome"`
	Marca string  `json:"Marca"`
	Valor float64 `json:"Valor"`
	Turbo bool    `json:"Turbo"`
}

func main() {

	Audi := Carros{"Audi A3", "Audi", 76.985, true}

	fmt.Println(Audi)

	mudeMe(&Audi)
	fmt.Println(Audi)

	// como alterei o valor com a func mudeMe ele irá transformar o valor alterado em json ( no caso o passat )
	audiJson, erro := json.Marshal(Audi)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(string(audiJson))

}

func mudeMe(c *Carros) {
	(*c).Nome = "Passat"
	c.Marca = "Volkswagen"
	c.Valor = 65.876
}
