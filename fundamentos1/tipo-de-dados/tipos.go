package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		idade          = 25
		firstname      = "Lucas"
		variavelPublic = false
		salario        = 5000.225
	)

	// type dos dados
	tpInt := reflect.TypeOf(idade)
	tpFloat64 := reflect.TypeOf(salario)
	tpStr := reflect.TypeOf(firstname)
	tpBool := reflect.TypeOf(variavelPublic)
	fmt.Println("Tipo int = ", tpInt)
	fmt.Println("Tipo float64 = ", tpFloat64)
	fmt.Println("Tipo string = ", tpStr)
	fmt.Println("Tipo bool = ", tpBool)
}
