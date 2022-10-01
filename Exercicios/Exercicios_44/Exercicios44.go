// Crie um struct de usuarios contendo os seguintes valores - nome, sobrenome, idade
// defina 3 usuarios para essa struct
// utilize marshal para transformar []user em JSON.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	Age      int    `json:"Age"`
}

func main() {

	User1 := User{"Gilberto", "Lopes", 65}

	User2 := User{"Joana", "Dark", 30}

	User3 := User{"Vitor", "Silva", 46}

	user1Json, err := json.Marshal(User1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(user1Json))

	user2Json, err := json.Marshal(User2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(user2Json))

	user3Json, err := json.Marshal(User3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(user3Json))

}
