package main

import "fmt"

type consoles struct {
	Marca   string
	Modelo  string
	Valor   int
	BlueRay bool
}

func (c consoles) salvar() {
	fmt.Printf("O console %v, no valor de R$ %v ,foi salvo no banco com exito", c.Modelo, c.Valor)
}

func (c consoles) desconto() {
	total := c.Valor * 10 / 100
	total2 := c.Valor - total
	fmt.Printf("\nO valor do %s é %v Reais, para pagamento a vista você terá um desconto de %v, saindo pelo valor de %v Reais", c.Modelo, c.Valor, total, total2)
}

func main() {

	PS5 := consoles{"Sony", "PS5", 5785, true}
	PS5.salvar()
	PS5.desconto()

	Nintendo := consoles{"Nintendo", "Switch", 3487, false}
	Nintendo.salvar()
	Nintendo.desconto()

}
