package main

import "fmt"

func imprimrResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Nota: ", nota, "Aprovado")
	} else {
		fmt.Println("Nota: ", nota, "Reprovado")
	}
}

func main() {
	nota := 5.0
	imprimrResultado(nota)
	nota = 8
	imprimrResultado(nota)
}
