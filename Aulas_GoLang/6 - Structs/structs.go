// Criei uma struct de Consoles, o qual declarei dos valores para a mesma, e transformei o json.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Consoles struct {
	Marca   string  `json:"Marca"`
	Modelo  string  `json:"Modelo"`
	Valor   float64 `json:"Valor"`
	BlueRay bool    `json:"Blue-ray"`
}

func main() {

	Sony := Consoles{
		Marca:   "Sony",
		Modelo:  "PS5",
		Valor:   5890.90,
		BlueRay: true,
	}
	fmt.Println("Struct de consoles")

	sonyJson, erro := json.Marshal(Sony)
	fmt.Println(string(sonyJson))
	if erro != nil {
		log.Fatal(erro)
	}

	Nintendo := Consoles{"Nintendo", "Wii-U", 1250.65, false}

	nintendoJson, erro := json.Marshal(Nintendo)
	fmt.Println(string(nintendoJson))
	if erro != nil {
		log.Fatal(erro)
	}
}
