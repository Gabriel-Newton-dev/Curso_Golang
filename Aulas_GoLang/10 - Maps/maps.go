package main

import "fmt"

func main() {

	marcaDeNotebooks := map[int]string{
		1: "Macbook",
		2: "Dell",
		3: "Acer",
		4: "Samsung",
		5: "Positivo",
	}
	fmt.Println(marcaDeNotebooks)

	fmt.Printf("Na minha opinião o melhor notebook é %v\n", marcaDeNotebooks[1])

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])
	delete(usuario, "sobrenome") // com delete excluímos aquele determinado sobrenome

	// irei fazer agora um um map de string que retorna um outra map
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Gabriel",
			"ultimo":   "Newton",
		},
		"curso": {
			"nome":   "Analise e Desenvolvimento de Sistemas",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Sagitário",
	}

	fmt.Println(usuario2)
}
