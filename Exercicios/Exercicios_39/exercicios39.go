// - Utilize a declaração defer de maneira que demonstre que sua
// execução só ocorre ao final do contexto ao qual ela pertence.

package main

import "fmt"

func main() {

	defer fmt.Println("Aqui apliquei o defer então será impresso por último")

	fmt.Println("Impressão 1")

	x := 2022

	fmt.Printf("Estamos no ano de %v\n", x)

}
