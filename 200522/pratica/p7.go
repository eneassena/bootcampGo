package pratica

import "math"

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Largura, Altura float64
}

type Geometria interface {
	Area() float64
	Perimentro() float64
}

func (t Triangulo) Area() float64 {
	return t.Largura * t.Altura
}

func (t Triangulo) Perimentro() float64 {
	return 2*t.Largura + 2*t.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimentro() float64 {
	return 2 * math.Pi * c.Raio
}

func NovoCirculo(values float64) Circulo {
	return Circulo{Raio: values}
}

const (
	tipoTriangulo = "triangulo"
	tipoCirculo   = "circulo"
)

func FiguraGeometrica(tipo string, values ...float64) Geometria {
	switch tipo {
	case tipoTriangulo:
		return Triangulo{Largura: values[0], Altura: values[1]}
	case tipoCirculo:
		return Circulo{Raio: values[0]}
	}
	return nil
}
