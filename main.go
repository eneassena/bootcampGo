package main

import (
	"errors"
	"fmt"

	"github.com/eneassena/bootcampGo/250522/modulo2"
)


 func main() {
	f := tarefas.Funcionario{Salario: 2000}
	
	// chamada de função de error basico  
	fmt.Println(fmt.Sprintln("Salario do funcionario = ", f.Salario))
	
	// função erro com o fmt.Errorf()
	msg, err := f.ErrorWithErrorf()
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println("Message: ", msg)
	

	// segundo error com a função errors.New()
	msg, err = f.ErrorWithNew()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Message: ", msg)
	
	// questão 4
	s, err := f.CalcSalarioHoras(44, 15)
	if err != nil {
		err1 := fmt.Errorf("error: %w", f)
		fmt.Println(errors.Unwrap(err1), err)
	}
	fmt.Println("salario: ",s)
}
