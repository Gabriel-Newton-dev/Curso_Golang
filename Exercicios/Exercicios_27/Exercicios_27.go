// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

package main

import "fmt"

func main() {

	x := struct {
		nome      string
		sabor     string
		ondeTem   []string
		vaiBemCom map[string]string
	}{
		nome:    "Sorvete",
		sabor:   "Napolitano",
		ondeTem: []string{"Sorveteria", "Padaria"},
		vaiBemCom: map[string]string{
			"No copo":      "Com cobertura",
			"Na casquinha": "Sem cobertura",
		},
	}
	fmt.Println(x)

}
