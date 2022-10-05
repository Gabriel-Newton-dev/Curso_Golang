// - Partindo do código abaixo, utilize marshal para transformar []user em JSON.
//     - https://play.golang.org/p/U0jea43X55
// - Atenção! Tem pegadinha aqui.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	First string
	Age   int
}

func main() {

	u1 := User{"James", 35}

	u2 := User{"Moneypenny", 27}

	u3 := User{"Maria", 54}

	Users := []User{u1, u2, u3}
	fmt.Println(Users)

	sb, err := json.Marshal(Users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(sb))
}
