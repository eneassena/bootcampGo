package exercicios

import "fmt"

func Exercicio5() {
	var texto string = "Hello, Word"

	tamanho := len(texto)

	fmt.Println("Tamanho da variavel", tamanho)

	for _, value := range texto {
		fmt.Printf("%c\n", value)
	}
}
