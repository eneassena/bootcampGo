package exercicios

/*
	Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
	notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
	de suas notas.

	Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
	máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
	definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
	indicado na função anterior
	Exemplo:
*/

const (
	MINIMO = "Menor"
	MAXIMO = "Maior"
	MEDIA  = "Media"
)

func minimo(numeros ...int) int {
	min := 0

	for k, num := range numeros {
		if k == 0 {
			min = num
		} else {
			if num < min {
				min = num
			}
		}
	}
	return min
}

func maximo(numeros ...int) int {
	max := 0

	for k, num := range numeros {
		if k == 0 {
			max = num
		} else {
			if num > max {
				max = num
			}
		}
	}
	return max
}

func media(numeros ...int) int {
	soma, media, total := 0, 0, 0

	for _, num := range numeros {
		soma += num
		total += 1
	}
	media = soma / total
	return media
}

func ExecFunctions(action string) func(numeros ...int) int {
	switch action {
	case MINIMO:
		return minimo
	case MAXIMO:
		return maximo
	case MEDIA:
		return media
	}
	return nil
}
