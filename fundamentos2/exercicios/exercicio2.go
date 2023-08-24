package exercicios

import (
	"errors"
	"fmt"
)

func ErrorsSystem() {
	num1 := 2
	var num2 = 0

	if num2 == 0 {
		err := errors.New("não é possivel dividir por zero")
		fmt.Println(err.Error())
	} else {
		fmt.Println(num1 / num2)
	}
}
