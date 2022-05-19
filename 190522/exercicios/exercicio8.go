package exercicios

import "fmt"

func Exercicio8() {

	var (
		employees        = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
		funcionario      = "Benjamin"
		countFuncionario = 0
	)

	for key, value := range employees {
		if key == funcionario {
			fmt.Printf("Funcionario %s foi encontrado\n", funcionario)
			fmt.Printf("%v\t%v\n", key, value)
		}
		if value > 21 {
			countFuncionario++
		}
	}

	fmt.Println("total de funcionarios com idade acima de 21 anos\n", countFuncionario)

	// - Saiba quantos de seus funcionários têm mais de 21 anos.
	fmt.Printf("Na empreza gaspazinho possui %v funcionários\n", len(employees))

	// - Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
	fmt.Println("O funciomário Federico foi Contratado")
	employees["Federico"] = 25

	// - Excluir Pedro do mapa.
	fmt.Println("O funcionário Pedro foi demitido")
	delete(employees, "Pedro")

	fmt.Printf("\n\n\n")
	fmt.Printf("==============================================\n")
	listarFuncionarios(employees)
}

func listarFuncionarios(employees map[string]int) {
	fmt.Printf("Vendo lista de funcionarios da empreza\n")
	for key, value := range employees {
		fmt.Printf("%v\t%v\n", key, value)
		fmt.Printf("======================\n")

	}
}
