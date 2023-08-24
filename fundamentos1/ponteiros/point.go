package main

import "fmt"

func main() {
	i := 1
	var p *int = nil // iniciando um pronteiro vasio
	p = &i
	fmt.Println(*p)
	*p++ // alterando o valor da variavel "i"
	fmt.Println(*p, i)
	fmt.Println(p == &i)

	var numero int = 10
	var ptr *int
	ptr = &numero
	fmt.Println(numero, *ptr)
}
