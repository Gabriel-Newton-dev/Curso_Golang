// Faça um map de , sendo key string e value [] string consoles videosgames antigos e novos, dizendo sobre os pontos positivos de cada console
// demostre os 2 primeiros e os 2 ultimos

package main

import "fmt"

func main() {
	// estrutura do map map[key]value{Colocar o nosso "dicionário aqui dentro"}
	consoles := map[string][]string{
		"PS4":             {"Sony", "Excelente jogos", "Blu-ray", "4K", "valor acessivel"},
		"PS5":             {"Sony", "Melhor vídeoGame dos últimos tempos", "Melhores jogos", "jogabilidade incrivel", "8k"},
		"XboxOne":         {"Microsoft", "Bom vide Game", "boas franquias", "boa jogabilidade", "valor acessivel"},
		"Nintendo Switch": {"Nintendo", "Muitos jogos infantis", "opção de ser portátil", "Jogos do Mario", "4k"},
	}
	for key, value := range consoles {
		fmt.Println(key, "-", value)
	}
}
