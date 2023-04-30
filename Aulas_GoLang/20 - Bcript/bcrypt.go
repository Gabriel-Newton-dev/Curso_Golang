package main

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 10
)

func main() {

	senha := "20julho1980"
	senhaErrada := "20julho1990"

	//func GenerateFromPassword(senha [] byte , custo int ) ([] byte , erro )
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), DefaultCost)
	if err != nil {
		errors.New("bcrypt: Error generating password", err)
	}
	fmt.Println(string(hashedPassword))

	//func CompareHashAndPassword(hashedPassword, senha [] byte ) erro
	//Nessa função iremos comparar o hashedPassword com a variavel senha
	fmt.Println(bcrypt.CompareHashAndPassword(hashedPassword, []byte(senha)))
	// //Nessa função iremos comparar o hashedPassword com a variavel senhaErrada
	fmt.Println(bcrypt.CompareHashAndPassword(hashedPassword, []byte(senhaErrada)))
}
