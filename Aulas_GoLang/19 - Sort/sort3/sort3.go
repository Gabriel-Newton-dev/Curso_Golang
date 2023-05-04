package main

import (
	"fmt"
	"sort"
)

func main() {

	slice := []int{10, 13, 14, 56, 11, 32, 15, 17, 12, 16, 18, 20, 21, 22, 19, 25}

	sort.Ints(slice)
	fmt.Printf("Colocando em ordem crescente o slice: %d\n", slice)

	sliceString := []string{"Maça", "Uva", "Goiaba", "Abobora", "Pessego"}

	sort.Strings(sliceString)
	fmt.Printf("Colocando em ordem alfabética o sliceString : %s\n", sliceString)
}
