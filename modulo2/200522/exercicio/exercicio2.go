package exercicio

import "errors"

func MediaAluno(numeros ...int) (int, error) {
	var sum = 0

	for _, v := range numeros {
		if v < 0 {
			return 0, errors.New("numero negativo impossivel somar")
		}
		sum += v
	}

	return sum / len(numeros), nil
}
