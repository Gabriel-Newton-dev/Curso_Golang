// Go não possui Herança, porém irei fazer um processo que é tipo Herança

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Carros struct {
	Marca  string  `json:"Marca"`
	Modelo string  `json:"Modelo"`
	Cor    string  `json:"Cor"`
	Turbo  bool    `json:"Turbo"`
	Valor  float64 `json:"Valor"`
}

type Caminhonete struct {
	Carros
	CabineDupla bool `json:"Cabine Dupla"`
	Importada   bool `json:"Importada"`
}

func main() {

	Cruze := Carros{"Chevrolet", "Cruze LTZ", "Prata", true, 99543.90}

	Ranger := Caminhonete{
		Carros:      Carros{"Ford", "Ranger", "Vinho", true, 120.690},
		CabineDupla: true,
		Importada:   false,
	}

	cruzeJson, erro := json.Marshal(Cruze)
	fmt.Println(string(cruzeJson))
	if erro != nil {
		log.Fatal(erro)
	}

	rangerJson, erro := json.Marshal(Ranger)
	fmt.Println(string(rangerJson))
	if erro != nil {
		log.Fatal(erro)
	}

}
