// Iremos utilizar o newEncoder e encode
// vamos usar method chaining para conectar os dois métodos acima.
// basicamente vc usa um retorno de um metodo, para utilizar em outro.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cell_Phones struct {
	Name        string
	Model       string
	Value       float64
	Description string
}

func main() {

	Iphones := Cell_Phones{"Iphone", "13 Pro Max", 7890.98, "Apple cell phone"}

	Samsung := Cell_Phones{"Samsung", "S-23", 6543.21, "Samsung cell phone"}

	Motorola := Cell_Phones{"Motorolo", "One", 2432.90, "Motorola cell phone"}

	CellPhones := []Cell_Phones{Iphones, Samsung, Motorola}

	//fmt.Println(CellPhones)

	err := json.NewEncoder(os.Stdout).Encode(CellPhones)
	if err != nil {
		fmt.Println(err)
	}
}
