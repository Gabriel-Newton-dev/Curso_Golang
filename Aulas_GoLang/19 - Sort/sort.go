package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {

	people := []Person{
		{"Gabriel", 37},
		{"Gisele", 39},
		{"Hulk", 40},
	}

	carros := []string{"Vectra", "Cruze", "Omega"}
	fmt.Println(carros)
	fmt.Println(people)

	sort.Strings(carros)
	fmt.Println(carros)

	// para fazer um sort em int

	si := []int{50, 70, 20, 10, 30, 40, 60}
	fmt.Println(si)
	sort.Ints(si) // para ordenar os inteiros do menor para o maior
	fmt.Println(si)

}
