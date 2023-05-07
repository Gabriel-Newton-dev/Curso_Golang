package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Carros struct {
	Name  string
	Model string
	Value float64
	Turbo bool
}

type Pickup_truck struct {
	Carros
	Imported    bool
	DoubleCabin bool
}

func main() {

	Cruze := Carros{"Cruze", "Ltz", 89.754, true}

	NewCivic := Carros{"New Civic", "Touring", 134.545, true}

	Silverado :=
		Pickup_truck{Carros: Carros{"Silverado", "Ltz", 784.563, true},
			Imported:    true,
			DoubleCabin: true,
		}

	fmt.Println(Silverado)

	AllCars := []Carros{Cruze, NewCivic}

	funcNewEncoder := json.NewEncoder(os.Stdout)

	err := funcNewEncoder.Encode(AllCars)
	if err != nil {
		fmt.Println(err)
	}

}
