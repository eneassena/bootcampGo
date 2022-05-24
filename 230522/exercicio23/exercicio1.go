package exercicio23

import "fmt"

type Produto struct {
	ID         int
	quantidade int
	preco      float64
}

func StartExercicio23() {
	p := Produto{ID: 1, quantidade: 15, preco: 35.00}

	fmt.Printf("ID do produto: %v\tQuantidade: %v\tpreço: R$ %.2f\n", p.ID, p.quantidade, p.preco)
}
