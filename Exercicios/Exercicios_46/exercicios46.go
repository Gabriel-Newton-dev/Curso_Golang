// - Partindo do código abaixo, utilize unmarshal e demonstre os valores.
//     - https://play.golang.org/p/b_UuCcZag9
// - Dica: JSON-to-Go.

package main

import (
	"encoding/json"
	"fmt"
)

type jsontogo []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var resultado jsontogo

	// formato do Unmarshall, declara o erro, chama a função json.Unmarshal recebendo um slice de byte do
	//json que no caso em tela é o s, e cria uma variavel para referenciar a struct que no caso é resultado, então usa um &

	err := json.Unmarshal([]byte(s), &resultado)
	if err != nil {
		fmt.Println("deu zica mano!", err)
	}

	fmt.Println(resultado)
	fmt.Println(resultado[1])
	fmt.Println(resultado[1].Last)

}
