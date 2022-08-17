// - Crie uma variável de tipo string utilizando uma raw string literal.
// - Demonstre-a.

package main

import "fmt"

func main() {

	x := `Criando uma variável       
	do tipo string
				utilizando o Raw
									STring` // para raw string usamos o `` em vez de utilizarmos aspas como de costume.
	fmt.Println(x)
}
