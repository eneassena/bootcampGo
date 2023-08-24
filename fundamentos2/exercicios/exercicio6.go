package exercicios

import "fmt"

// A denifição de uma estrutura se da pela: determinação de seus campos seguidos por um espaço e o tip de dados
// para separa os campos usa-se uma quebra de linha

type Veiculo struct {
	km    float64
	horas float64
}

type Automovel struct {
	v Veiculo
}

type Moto struct {
	v Veiculo
}

func (v Veiculo) detalhes() {
	fmt.Printf("\nKM: %.2f\tHoras: %.2f\n", v.km, v.horas)
}

func (a *Automovel) Correr(minutos int) {
	a.v.horas = float64(minutos) / 100
	a.v.km = a.v.horas * 100
}

func (a *Automovel) Detalhes() {
	fmt.Println("\nV: Automovel")
	fmt.Printf("KM: %.2f\t Horas: %.2f", a.v.horas, a.v.km)
	a.v.detalhes()
}

func (a *Moto) Correr(minutos int) {
	a.v.horas = float64(minutos) / 100
	a.v.km = a.v.horas * 80
}

func (a *Moto) Detalhes() {
	fmt.Println("\nV: Moto")
	fmt.Printf("\nKM: %.2f\tHoras: %.2f", a.v.km, a.v.horas)
	a.v.detalhes()
}
