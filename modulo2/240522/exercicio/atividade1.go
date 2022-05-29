package base03

import (
	"errors"
	"fmt"
)

type ErrorPersonalizado interface {
	Error() string
}

const minimoSalario = 15000

type Funcionario struct {
	Salario int
}

const messageError string = "[Error]: Necessário pagar imposto"

// Questão 1
func (f Funcionario) Error() string {
	if f.Salario < minimoSalario {
		return fmt.Sprintf("[Error] error: Salario digitado nao esta dentro do valor minimo que é = R$: %v",
			minimoSalario)
	}
	fmt.Println(messageError)
	return ""
}

// Questão 2
func (funcionario Funcionario) ErrorWithNew() (string, error) {
	if funcionario.Salario < minimoSalario {
		m := fmt.Sprintf("[ErrorWithNew] error: func New: Salario digitado nao esta dentro do valor minimo que é = R$: %d",
			minimoSalario)
		return "", errors.New(m)
	}
	return messageError, nil
}

// Questão 3
func (funcionario Funcionario) ErrorWithErrorf() (string, error) {
	if funcionario.Salario < minimoSalario {
		messageErr := fmt.Errorf("[ErrorWithErrorf] error: Salario digitado nao esta dentro do valor minimo que é = R$: %d",
			minimoSalario)
		return "", messageErr
	}
	return messageError, nil
}

// questão 4
func (funcionario Funcionario) CalcSalarioHoras(horas int, valor int) (int, error) {
	if horas < 80 {
		err := fmt.Errorf("error: o trabalhador não pode ter trabalhado menos de %d horas por mês", horas)
		return 0, err
	}
	funcionario.Salario = valor * horas

	if funcionario.Salario >= 20000 {
		// desconto de 10%
		imposto := funcionario.Salario * (10 / 100)
		funcionario.Salario -= imposto
	}
	return funcionario.Salario, nil
}
