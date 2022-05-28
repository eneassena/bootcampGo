package base04

import "fmt"

func debug() {
	err := recover()

	if err != nil {
		fmt.Println("error: ", err)
	}
}

func isEnv(num int) {
	if num%2 != 0 {
		//defer debug()
		panic("Nao é um numero par")
	}
	fmt.Println(num, " é um número par")
}
