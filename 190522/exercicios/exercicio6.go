package exercicios

import "fmt"

func Exercicio6() {
	var (
		idade     int
		salario   int
		atividade int
		empregado bool
	)

	idade, salario, atividade, empregado = 25, 100001, 10, true

	var message string

	if idade > 22 && (empregado && atividade > 12) {
		// cliente aprovado para empréstimo
		message = "Cliente apto para empréstimo"

		if salario <= 100000 {
			// cliente aprovado para empréstimo, mais pagará juros
			message = "Cliente apto para empréstimo, Com jutos"
		}
	} else {
		// cliente nao foi aprovado para empréstimo
		message = "Cliente não foi aprovado para empréstimo"
	}

	fmt.Println(message)

	/*switch {
	case idade > 22 && atividade > 12 && empregado && salario > 100000:
		fmt.Println("Cliente Apto a pega empréstimo sem juros")
	case idade > 22 && atividade > 12 && empregado && salario < 100000:
		fmt.Println("Cliente Apto a receber empréstimo com Juros")
	default:
		fmt.Println("Cliente Não pode está apto a receber empréstimo")
	}*/
}
