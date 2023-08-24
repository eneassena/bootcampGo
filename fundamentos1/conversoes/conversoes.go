package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.5
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	fmt.Println(nota)
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado
	fmt.Println("Test " + string(132))

	fmt.Println("Test " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("324")
	fmt.Println(num + 1)

	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}
}
