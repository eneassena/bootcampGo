package exercicios

import (
	"fmt"
)

const (
	A = "A"
	B = "B"
	C = "C"
)

func SalarioHoras() {
	var (
		categoria      string
		horasAtividade float64
		salario        float64
	)

	categoria = C
	horasAtividade = 162.0

	switch categoria {
	case A:
		salario = 3000.0 * horasAtividade

		if horasAtividade > 160 {
			salario += salario * 0.50
		}
	case B:
		salario = 1500.00 * horasAtividade

		if horasAtividade > 160 {
			salario += salario * 0.20
		}
	case C:
		salario = 1000 * horasAtividade

	}
	fmt.Println(categoria, horasAtividade, salario)
}
