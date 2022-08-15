// - Utilizando a solução do exercício anterior:
//     1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com
// 	o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
//     2. Na função main:
//         1. Isto já deve estar feito:
//             1. Demonstre o valor da variável "x"
//             2. Demonstre o tipo da variável "x"
//             3. Atribua 42 à variável "x" utilizando o operador "="
//             4. Demonstre o valor da variável "x"
//         2. Agora faça tambem:
//             1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
//             2. Demonstre o valor de "y"
//             3. Demonstre o tipo de "y"

package main

import "fmt"

type number int

var x number
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("O tipo de X é %T\n", x)
	x = 42
	fmt.Println(x)

	y := x
	fmt.Println(y)
	fmt.Printf("O tipo de Y é %T", y)
}
