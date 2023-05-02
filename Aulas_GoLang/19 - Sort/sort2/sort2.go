package main

import (
	"fmt"
	"sort"
)

func main() {

	sliceInt := []int{9, 8, 4, 3, 5, 1, 10, 31, 56, 6, 7, 98, 10}

	sort.Ints(sliceInt)
	fmt.Println(sliceInt)

	sliceString := []string{"Zebra", "Leão", "Avestruz", "Macaco", "Papagaio", "Cachorro", "Elefante", "Girafa", "Sapo", "Hipopotámo"}

	sort.Strings(sliceString)
	fmt.Println(sliceString)
}
