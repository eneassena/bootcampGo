package main

import (
	"fmt"

	"github.com/eneassena/bootcampGo/200522/exercicio"
)

const (
	GATO      = "gato"
	CAES      = "caes"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

func main() {

	// Exercicio 4
	minimo := exercicio.ExecFunctions("Menor")
	maximo := exercicio.ExecFunctions("Maior")
	media := exercicio.ExecFunctions("Media")

	fmt.Println("Menor dos valor da entra: [10, 4, 5, 6, 7, 8, 9, 56, 4, 34, -45, 2, 3]")
	fmt.Println(minimo(10, 4, 5, 6, 7, 8, 9, 56, 4, 34, -45, 2, 3))

	fmt.Println("\nMaior dos valor da entra: [10, 4, 5, 6, 7, 8, 9, 56, 4, 34, -45, 2, 3]")
	fmt.Println(maximo(10, 4, 5, 6, 7, 8, 9, 56, 4, 34, -45, 2, 3))

	fmt.Println("\nMédia dos valor da entra: [3, 3, 3]")
	fmt.Println(media(3, 3, 3))

	// Exercicio 5
	var amount int

	// animalGato, msgGato := exercicio.Animal("chani") // retorna uma mensagem de erro
	animalGato, msgGato := exercicio.Animal(GATO) // retorna a função do animal Gato corretamente
	animalCaes, msgCaes := exercicio.Animal(CAES)
	animalHamster, msgHamster := exercicio.Animal(HAMSTER)
	animalTarantula, msgTarantula := exercicio.Animal(TARANTULA)

	if animalGato != nil {
		amount += animalGato(5)
	}
	amount += animalCaes(3)
	amount += animalHamster(8)
	amount += animalTarantula(4)

	if msgCaes != nil || msgGato != nil || msgHamster != nil || msgTarantula != nil {
		fmt.Println("Mensagem pro animal Gato", msgGato)
		fmt.Println("Mensagem pro animal Gato", msgCaes)
		fmt.Println("Mensagem pro animal Hamster", msgHamster)
		fmt.Println("Mensagem pro animal Tarantula", msgTarantula)
	}

	fmt.Println("o total de alimento será de = ", amount)

}
