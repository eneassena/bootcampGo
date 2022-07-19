package exercicio

func FuncionarioSalario(salarioBruto float64) float64 {
	var (
		juros   float64
		imposto float64
	)
	if salarioBruto <= 50000 {
		juros = 17
	} else if salarioBruto <= 150000 {
		juros = 27
	}
	imposto = salarioBruto * (juros / 100)
	return imposto
}
