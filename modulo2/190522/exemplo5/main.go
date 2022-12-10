package main

import "fmt"

func main() {
	var texto string = "Hello, Word"

	tamanho := len(texto)

	fmt.Println("Tamanho da variavel", tamanho)

	for _, value := range texto {
		fmt.Printf("%c\n", value)
	}
}
