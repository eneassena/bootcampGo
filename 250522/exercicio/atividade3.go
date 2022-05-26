package exercicio

import (
	"errors"
)

// implementacao de funcao de erro evolucao
func (v VarTypeErro) ErrorTwo() (string, ErrorPersonalizado) {
	if v.Salario < 15000 {
		return "", errors.New("error: Salario digitado nao esta dentro do valor minimo")
	}
	// fmt.Println("Necessário pagar imposto")
	return "Necessário pagar imposto", nil
}
