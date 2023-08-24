package exercicios

import "errors"

func MediaAluno(numeros ...int) (int, error) {
	var (
		sum        = 0
		media      = 0
		totalNotas = len(numeros)
	)

	for _, v := range numeros {
		if v < 0 {
			return 0, errors.New("número negativo impossivel Calcular Média do aluno")
		}
		sum += v
	}
	media = sum / totalNotas
	return media, nil
}
