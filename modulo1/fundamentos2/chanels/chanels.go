package main

import "fmt"

func ehMaiorDeIdade(idade int, c chan bool) {
	ehMaior := idade > 17
	c <- ehMaior
}

func main() {
	c := make(chan bool)

	idade := 18
	go ehMaiorDeIdade(idade, c)
	ehMaior := <-c

	fmt.Println("O aluno tem,", idade, "anos")
	if ehMaior {
		fmt.Println("e com essa dade ele é maior de idade")
	} else {
		fmt.Println("e com essa dade ele é menor de idade")
	}

}
