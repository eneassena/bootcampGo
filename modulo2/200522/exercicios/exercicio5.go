package exercicios

import "errors"

/* Descrição da atividade 5

Um abrigo de animais precisa descobrir quanta comida comprar para os animais de
estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão
é que haja muito mais animais para abrigar.
1. Cães precisam de 10 kg de alimento
2. Gatos 5 kg
3. Hamster 250 gramas.
4. Tarântula 150 gramas.

Precisamos:
1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto
com o animal especificado e que retorne uma função com uma mensagem (caso não
exista o animal)
2. Uma função para cada animal que calcule a quantidade de alimento com base na
quantidade necessária do animal digitado.
Exemplo:

var (
	caes = 10 // kg
	gato = 5 //kg
	hamster = 250 // gramas
	tarantula = 150 // gramas
)

função tem um paramentro do tipo string, retorna uma mensagem de erro caso o animal nao exite

const (
	dog = "dog"
	cat = "cat"
)


animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)


var amount float64
amount+= animaldog(5)
amount+= animalCat(8)
*/

const (
	CAES      = 10  // kg
	GATO      = 5   //kg
	HAMSTER   = 250 // gramas
	TARANTULA = 150 // gramas
)

func gato(quantidade int) int {
	return GATO * quantidade
}
func caes(quantidade int) int {
	return CAES * quantidade
}
func hamster(quantidade int) int {
	kg := (HAMSTER * 4) / 1000
	return kg * quantidade
}
func tarantula(quantidade int) int {
	kg := (TARANTULA * 8) / 1000
	return kg * quantidade
}

func Animal(animal string) (func(quantidade int) int, error) {
	if animal == "gato" {
		return gato, nil
	} else if animal == "caes" {
		return caes, nil
	} else if animal == "hamster" {
		return hamster, nil
	} else {
		if animal == "tarantula" {
			return tarantula, nil
		}
	}
	msg := errors.New("Animal não foi encontrado :(")
	return nil, msg
}
