// - Partindo do código abaixo, ordene a []int e a []string.
//     - https://play.golang.org/p/H_q75mpmHW

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 10, 17, 9, 14, 7, 12, 21, 1, 4, 6, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
}
