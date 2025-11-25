package main

/*
	1.2 -> suponha que você duplique o tamanho da lista. Qual seria o numero maximo de etapas agora?

	log2(256) = 8
	r: 8 tentativas
*/

import (
	"fmt"
	"sort"
)

func main() {
	nomes := []string{
		// 1-32
		"Adriana", "Bruno", "Carlos", "Daniela", "Eduardo", "Fernanda", "Gabriel", "Helena",
		"Igor", "Julia", "Kleber", "Larissa", "Marcos", "Natalia", "Otavio", "Paula",
		"Quintino", "Rafael", "Sabrina", "Thiago", "Ursula", "Vinicius", "Wagner", "Xavier",
		"Yasmin", "Zeca", "Alice", "Bernardo", "Caio", "Diana", "Elias", "Fabio",
		// 33-64
		"Gustavo", "Heitor", "Isabela", "Joana", "Karen", "Lucas", "Matheus", "Nicole",
		"Olivia", "Pedro", "Quezia", "Renato", "Sofia", "Tatiana", "Ubirajara", "Vitoria",
		"Wesley", "Ximena", "Yuri", "Zaira", "Amanda", "Breno", "Camila", "Diego",
		"Elisa", "Felipe", "Gabriela", "Henrique", "Isaac", "Jessica", "Kevin", "Livia",
		// 65-96
		"Miguel", "Nair", "Orlando", "Patricia", "Queiroga", "Ricardo", "Sandra", "Tomas",
		"Ulisses", "Vanessa", "William", "Xandao", "Yara", "Zuleica", "Antonio", "Bianca",
		"Cesar", "Debora", "Emilio", "Fatima", "Gilberto", "Hilda", "Ines", "Jorge",
		"Kelly", "Leandro", "Monica", "Nelson", "Olavo", "Pamela", "Quiteria", "Roberto",
		// 97-128
		"Silvia", "Tereza", "Ubaldo", "Valter", "Wellington", "Xuxa", "Yago", "Zilma",
		"Alex", "Barbara", "Cristiano", "Doralice", "Evandro", "Flavia", "Gerson", "Hugo",
		"Ivana", "Jonas", "Katia", "Luciano", "Marina", "Norberto", "Odete", "Pericles",
		// 129-160
		"Abel", "Betina", "Clarice", "Davi", "Ester", "Fabricio", "Giovana", "Heraldo",
		"Iara", "Jefferson", "Karina", "Lorena", "Marcelo", "Noemia", "Osvaldo", "Pilar",
		"Quenia", "Rodrigo", "Simone", "Tulio", "Ugo", "Veronica", "Wallace", "Xenia",
		"Yone", "Zeneide", "Arthur", "Beatriz", "Claudio", "Douglas", "Eliane", "Francisco",
		// 161-192
		"Gisele", "Hamilton", "Ingrid", "Jaqueline", "Kleber", "Leonardo", "Manuela", "Nilson",
		"Otavia", "Priscila", "Queren", "Raimundo", "Sueli", "Tania", "Uriel", "Viviane",
		"Washington", "Xisto", "Yasmim", "Zilda", "Andre", "Brenda", "Cecilia", "Danilo",
		"Evelyn", "Fernando", "Guilherme", "Heloisa", "Italo", "Juliana", "Kamila", "Luana",
		// 193-224
		"Mario", "Nubia", "Osorio", "Paloma", "Querubina", "Rogerio", "Sheila", "Talita",
		"Ulisses", "Vania", "Waldir", "Xerxes", "Yuri", "Zaqueu", "Augusto", "Benedito",
		"Cintia", "Diogo", "Erica", "Fausto", "Glaucia", "Hermes", "Ivan", "Janaina",
		"Kaua", "Leticia", "Murilo", "Neide", "Oscar", "Poliana", "Quintiliano", "Renan",
		// 225-256
		"Sara", "Teodoro", "Uilian", "Valentina", "Wilma", "Xisto", "Yan", "Zelia",
		"Aline", "Bruna", "Carolina", "Daniel", "Estela", "Frederico", "Gustav", "Helio",
		"Iris", "Joaquim", "Kauane", "Liz", "Melissa", "Nicolas", "Odorico", "Pietra",
		"Queli", "Romulo", "Samanta", "Theo", "Ula", "Vicente", "Walter", "Zoraide",
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
