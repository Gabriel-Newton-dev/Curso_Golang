// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.

package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func (p Pessoa) completo() {
	fmt.Println(p.Nome, p.Sobrenome)
}

func main() {

	carlos := Pessoa{"Carlos Alberto", "de Nobrega", 85}

	Pessoa.completo(carlos)
	// para chamar o metodo é dessa forma, coloca o nome da struct.completo
	//(parametro que passei no metodo ) e entre parenteses o tipo da struct pessoa que criei

}
