package exercicios

import "fmt"

func Exercicio8() {
	/*
		*** Valores da flag na chamada da função: listarFuncionarios ***

		1: Buscar pelo funcionario: Benjamin
		2: Lista funcionarios com idade acima de 21 anos
		3: Listar todos funcionario
	*/

	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30}

	// - Saiba quantos de seus funcionários têm mais de 21 anos.
	fmt.Printf("Total de funcionários é = %v \n\n", len(employees))

	// Buscar pelo funcionario: Benjamin
	// fmt.Println("\n")
	listarFuncionarios(employees, 1)

	// Lista funcionarios com idade acima de 21 anos
	fmt.Println("\nlistagem de funcionarios com idade acima de 21 anos")
	listarFuncionarios(employees, 3)

	// - Excluir Pedro do mapa.
	fmt.Println("\nO funcionário Pedro foi demitido")
	delete(employees, "Pedro")

	// - Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
	fmt.Println("\nO funciomário Federico foi Contratado")
	employees["Federico"] = 25

	// Listar todos funcionario
	fmt.Println("\nListagem de Funcionarios")
	listarFuncionarios(employees, 2)

}

func listarFuncionarios(employees map[string]int, flag int) {

	funcionario := "Benjamin"

	for key, value := range employees {
		switch flag {
		case 1:
			if key == funcionario {
				fmt.Printf("Funcionario %s foi encontrado\n", funcionario)
				showData(key, value)
			}
		case 2:
			showData(key, value)
		case 3:
			if value > 21 {
				showData(key, value)
			}
		}
	}
}

func showData(key string, value int) {
	fmt.Printf("Nome: %v\tIdade: %v\n", key, value)
	fmt.Printf("===========================\n")
}
