package mapping

import "fmt"

func Mapeamento1() {
	studentes := map[string]int{"joao": 20, "Pedro": 26}
	fmt.Println(studentes["joao"])
	studentes["Ana"] = 35
	studentes["Maria"] = 42
	fmt.Println(studentes["joao"])

	alunos := make(map[string]interface{})
	alunos["Nome"] = "Pedro"
	alunos["Idade"] = 30
	fmt.Println(alunos)

	delete(alunos, "Idade")
	fmt.Println(alunos)
}
