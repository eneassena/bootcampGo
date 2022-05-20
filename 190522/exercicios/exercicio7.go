package exercicios

import "fmt"

func Exercicio7() {
	var (
		mes     int    = 5
		message string = "Mês Desconhecido"
	)

	switch mes {
	case 1:
		message = "Janeiro"
	case 2:
		message = "Fevereiro"
	case 3:
		message = "Março"
	case 4:
		message = "Abril"
	case 5:
		message = "Maio"
	case 6:
		message = "Junho"
	case 7:
		message = "Julho"
	case 8:
		message = "Agosto"
	case 9:
		message = "Setembro"
	case 10:
		message = "Outubro"
	case 11:
		message = "Novembro"
	case 12:
		message = "Dezembro"
	}

	fmt.Println(message)
}
