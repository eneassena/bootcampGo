package main

import (
	"fmt"
)

const TITLE string = "Declaração de varaveis"

var (
	descrição string = "exemplo de declaração de variaveis no go lang"
)

var idade int = 25

// criando contante
const pi float32 = 3.1415

const (
	idade1 = 15
	idade2 = 35
	idade3 = 14
)

func main() {
	fmt.Println(TITLE)

	fmt.Println(descrição)

	// declaração curta
	nome := "eneas"
	filme := "velozes e furiosos"
	fmt.Println(nome)
	fmt.Println(filme)

	anoLancamento := 2018
	fmt.Println(anoLancamento)

	VariavelPublic := true
	variavelPublic := false
	fmt.Printf("\nStatus para escopo da variavel publica = %t", VariavelPublic)
	fmt.Printf("\nStatus para escopo da variavel privada = %t", variavelPublic)

	sobrenome := "Sena"

	// fmt.Println("Hello World")
	var firstnome string

	firstnome = "eneas"

	fmt.Println()
	fmt.Println(firstnome, sobrenome)
	fmt.Println(len(firstnome))

	fmt.Println(firstnome, idade, pi)

	fmt.Println("\nconstantes")
	fmt.Println(idade1, idade2, idade3)

	var salario float64 = 50000.599
	fmt.Printf("salário, R$%.2f \n", salario)

	char := 'b'
	fmt.Println(char)
}
