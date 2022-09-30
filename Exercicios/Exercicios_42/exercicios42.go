// Faça um programa que ordene a strings por ordem alfabética

package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Package Sort serve para ordenar os itens de um slices int ou string")

	sliceString := []string{"Fabiana", "Luana", "Gilberto", "Ana", "Diego", "Carolina", "Ellen", "Bianca", "Zeus", "Joana"}

	sort.Strings(sliceString)
	fmt.Println(sliceString)

	sliceInt := []int{9, 4, 5, 2, 1, 3, 6, 10, 8, 11, 15, 13, 14, 12}

	sort.Ints(sliceInt)
	fmt.Println(sliceInt)

}
