package main

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	Nome          string `json:"nome"`
	Idade         int    `json:"idade"`
	Especialidade string `json:"especialidade"`
}

func main() {

	jsonStr := `[{"Nome":"Raphael","Idade":19,"Especialidade":"JavaScript"},{"Nome":"Gabriel","Idade":24,"Especialidade":"GoLang"},{"Nome":"Marco","Idade":23,"Especialidade":"BitCoin"}]`
	fmt.Println(jsonStr)
	var s []Users
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

}
