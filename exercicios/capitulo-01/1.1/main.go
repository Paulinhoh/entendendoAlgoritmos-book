package main

/*
	1.1 -> suponha que você tenha uma lista com 128 nomes e esteja fazendo uma pesquisa binaria. Qual seria o numero maximo de etapas que você levaria para encontrar o nome desejado?

	log2(128) = 7
	r: 7 tentativas
*/

import (
	"fmt"
	"sort"
)

func main() {
	nomes := []string{
		"Adriana", "Bruno", "Carlos", "Daniela", "Eduardo", "Fernanda", "Gabriel", "Helena",
		"Igor", "Julia", "Kleber", "Larissa", "Marcos", "Natalia", "Otavio", "Paula",
		"Quintino", "Rafael", "Sabrina", "Thiago", "Ursula", "Vinicius", "Wagner", "Xavier",
		"Yasmin", "Zeca", "Alice", "Bernardo", "Caio", "Diana", "Elias", "Fabio",
		"Gustavo", "Heitor", "Isabela", "Joana", "Karen", "Lucas", "Matheus", "Nicole",
		"Olivia", "Pedro", "Quezia", "Renato", "Sofia", "Tatiana", "Ubirajara", "Vitoria",
		"Wesley", "Ximena", "Yuri", "Zaira", "Amanda", "Breno", "Camila", "Diego",
		"Elisa", "Felipe", "Gabriela", "Henrique", "Isaac", "Jessica", "Kevin", "Livia",
		"Miguel", "Nair", "Orlando", "Patricia", "Queiroga", "Ricardo", "Sandra", "Tomas",
		"Ulisses", "Vanessa", "William", "Xandao", "Yara", "Zuleica", "Antonio", "Bianca",
		"Cesar", "Debora", "Emilio", "Fatima", "Gilberto", "Hilda", "Ines", "Jorge",
		"Kelly", "Leandro", "Monica", "Nelson", "Olavo", "Pamela", "Quiteria", "Roberto",
		"Silvia", "Tereza", "Ubaldo", "Valter", "Wellington", "Xuxa", "Yago", "Zilma",
		"Alex", "Barbara", "Cristiano", "Doralice", "Evandro", "Flavia", "Gerson", "Hugo",
		"Ivana", "Jonas", "Katia", "Luciano", "Marina", "Norberto", "Odete", "Pericles",
	}
	sort.Strings(nomes)

	buscaBinaria(nomes, nomes[len(nomes)-1])
}

func buscaBinaria(lista []string, alvo string) {
	inicio := 0
	final := len(lista) - 1
	tentativasFeitas := 0

	for inicio <= final {
		meio := (final + inicio) / 2
		chute := lista[meio]

		tentativasFeitas++
		if chute == alvo {
			fmt.Println("nome encontrado na lista:", chute)
			break
		}
		if chute > alvo {
			final = meio - 1
		}
		if chute < alvo {
			inicio = meio + 1
		}
	}

	fmt.Println("Tentativas feitas:", tentativasFeitas)
}
