package main

import (
	"fmt"
	"github.com/eneassena/bootcampGo/250522/exercicio"
)

func main() {
	salario := exercicio.VarTypeErro{Salario: 18000}
	ok, err := salario.ErrorTwo()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("saida de Error: ", ok)
	fmt.Println(salario.Error())

}
