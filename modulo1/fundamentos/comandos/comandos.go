package main

import (
	"fmt"
)

/**
acessar documentação do golang ofiline
comando:
- godoc -http=:6060
*/

func cincoVezes(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
}

func main() {
	c := make(chan int)
	for i := 1; i <= 5; i++ {
		go cincoVezes(c)
		res := <-c
		fmt.Println(res)
	}
	//fmt.Printf("Outro Programa em %s")
}
