// - O sort que eu quero não existe. Quero fazer o meu.
// - Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
//     - type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
// - Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
// - E aí posso fazer do jeito que eu quiser.
// - Exemplo:
//     - struct carros: nome, consumo, potencia
//     - slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
//     - tipo ordenarPorPotencia
//     - tipo ordenarPorConsumo
//     - tipo ordenarPorLucroProDonoDoPosto

package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

// interface slice carro
type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int           { return len(x) }
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int           { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
func (a ordenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int           { return len(x) }
func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	carros := []carro{
		{"Chevete", 50, 8},
		{"Porsche", 300, 5},
		{"Fusca", 20, 30},
	}

	fmt.Println("Inicial:\n", carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Potência:\n", carros)

	sort.Sort(ordenarPorConsumo(carros))

	fmt.Println("Consumo:\n", carros)

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))

	fmt.Println("Lucro:\n", carros)

}
