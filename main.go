package main

import (
	"fmt"
	"testing"
)


func Soma(a, b int) int {
	return a + b
}

func TestSoma(t *testing.T){
	result := Soma(5,5)
	esperado := 10

	if result != esperado {
		t.Errorf("resultado incorrento espera %d e a função Soma retorno %d", esperado, result)
	}


}


func main() {  
	fmt.Println("rodando testes")
}