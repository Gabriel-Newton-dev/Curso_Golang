// - Faça um programa que verifique se o alunho foi aprovado usando else if e else., lemrando que a nota 7.0

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Alunos struct {
	Nome  string
	Serie string
	Nota  float64
}

func main() {

	Caio := Alunos{
		Nome:  "Caio Santos",
		Serie: "5 Série",
		Nota:  8.9}

	Alunos := (Caio.Nota)
	if media := 7.0; Alunos >= media {
		fmt.Println("Aluno Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

	CaioJson, err := json.Marshal(Caio)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(CaioJson))
}
