package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	first string
	age   int
}

func main() {
	u1 := user{
		first: "James",
		age:   32,
	}

	u2 := user{
		first: "Moneypenny",
		age:   27,
	}

	u3 := user{
		first: "M",
		age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	user, err := json.Marshal(users)
	str1 := bytes.NewBuffer(user).String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("String =", str1)

}
