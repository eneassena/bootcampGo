package pratica

func opSoma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func Subtra(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func opMult(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func opDiv(valor1, valor2 float64) float64 {
	return valor1 / valor2
}

func Operator(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case "Soma":
		return opSoma
	case "Subtra":
		return Subtra
	case "Mult":
		return opMult
	case "Div":
		return opDiv
	}
	return nil
}
