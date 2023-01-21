package exemplos

import (
	"fmt"
)

func Exemplo1(nomes ...string) {
	fmt.Println("total de nomes recebidos", len(nomes))
}

func FullName(firstName string, lastName string) {
	fmt.Println(firstName + lastName)
}

func Exemplo2(firstName string, lastName string, action func(firstName string, lastName string)) {
	action(firstName, lastName)
}
