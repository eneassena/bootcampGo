package pratica

import (
	"errors"
	"fmt"
)

func VariavelFunc(a, b, c float64) {
	fmt.Printf("\n%v\t%v\t%v\n", a, b, c)
}

func AumuladorSoma(numbers ...float64) float64 {
	var acumulador float64
	for _, v := range numbers {
		acumulador += v
	}
	return acumulador
}

func RetornoMultiplo(valor1 int, valor2 int) (int, int, int, error) {
	sum := valor1 + valor2
	mult := valor1 * valor2
	if valor2 == 0 {
		return 0, 0, 0, errors.New("Não é possivel a divisão por zero 0")
	}
	div := valor1 / valor2


	return sum, mult, div, nil
}
