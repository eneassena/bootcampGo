package exemplos

import "fmt"

func debug() {
	if err := recover(); err != nil {
		fmt.Println("error: ", err)
	}
}

func isEnv(num int) {
	if num%2 != 0 {
		defer debug()
		panic("Nao é um numero par")
	}
	fmt.Println(num, " é um número par")
}

func Execute() {
	isEnv(5)
}
