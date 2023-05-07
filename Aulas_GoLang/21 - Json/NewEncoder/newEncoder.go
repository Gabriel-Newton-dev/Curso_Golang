// - Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.
//     - https://play.golang.org/p/BVRZTdlUZ_

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// precisamos de um NewEncoder para ai sim usar o encoder

	// funcNewEncoder := json.NewEncoder(os.Stdout)

	// err := funcNewEncoder.Encode(users)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// - Desafio: descubra o que é, e utilize, method chaining para conectar os dois métodos acima.
	// basicamente vc usa um retorno de um metodo, para utilizar em outro, ex:

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}

}
