package exercicio

import "fmt"

type ErrorPersonalizado interface {
	Error() string
}

type VarTypeErro struct {
	Salario int
}

func (v VarTypeErro) Error() string {
	if v.Salario < 15000 {
		return "error: Salario digitado nao esta dentro do valor minimo"
	}
	fmt.Println("Necessário pagar imposto")
	return ""
}
