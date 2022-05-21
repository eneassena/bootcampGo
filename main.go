package main

import (
	"fmt"

	"github.com/eneassena/bootcampGo/200522/pratica"
)

func main() {

	// fmt.Println("hello, World! :)")

	Subtra := pratica.Operator("Subtra")
	fmt.Println(Subtra(10, 3))

	// s, m, d, err := pratica.RetornoMultiplo(5, 5)
	// fmt.Printf("a soma eh = %v\te a multiplicação = %v\t Divisão = %v\tError = %v\n", s, m, d, err)

	// res := pratica.AumuladorSoma(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// fmt.Println(res)

	// stundents := map[string]int{"caio": 32, "Brenda": 12}

	// for k, v := range stundents {
	// 	fmt.Printf("%v - %v\n", k, v)
	// }

	// myMap := map[string]int{}
	// myMap := make(map[string]int)

	// fmt.Println(len(myMap))

	// var nomes [2]string
	// var idades [2]int

	// var nomes2 = make([]string, 2)

	// nomes[0] = "Eneas"
	// nomes[1] = "Mathues"
	// nomes2[0] = nomes[0]
	// nomes2[1] = nomes[1]

	// idades[0] = 35
	// idades[1] = 33

	// fmt.Printf("aluno 1: %s %v\n", nomes2[0], idades[0])
	// fmt.Printf("aluno 2: %s %v\n", nomes2[1], idades[1])

	// a := make([]int, 5)
	// a = append(a, 85)

	// fmt.Println(len(a), a[5])
	// var status = []bool{true, false}
	// fmt.Println(status[0])
	// fmt.Println(status[1])

	// var alunos = make(map[string]int{"Eneas": 35, "Matheus": 33}, 2)
	// fmt.Println(alunos)

}
