package pratica

import (
	"fmt"
	"os"
	"time"
)

/**
 * Pratica do pacote os
 *
 * func exploradas aqui
 * - os.Setenv(key, value string)
 *
 */

// func Setenv
func exemploOs1() {
	err := os.Setenv("Nome", "GOLANG")
	if err != nil {
		fmt.Println(err)
	}
}

// func Getenv
func exemploOs2() {
	value := os.Getenv("Nome")
	fmt.Println(value)

}

// func LookupEnv
func exemploOs3() {
	value, ok := os.LookupEnv("Nome")
	fmt.Println("OK: ", ok)
	fmt.Println("value: ", value)
}

// usando ponteiros/ forma de criação
func ponteirosVar() {
	var p1 *int
	var p2 = new(int)
	v := 10
	p2 = &v

	fmt.Println("ponteiro vasio = ", p1)
	fmt.Println("ponteiro criado em p2 = ", p2)
	fmt.Println("valor do ponteiro p2 = ", *p2)
	fmt.Println("endereço de memoria de v = ", &v)
}

func aumentar(v *int) {
	// acessando o valor na memoria e incrementando + 1
	*v++
}

func processar() {
	fmt.Println("-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("-Termina")
}

func processar2(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
	c <- i
}

func criandoCanal() {
	c := make(chan int)
	fmt.Println(c)
}

func PlayOs() {
	// execução normal
	// for i := 0; i < 10; i++ {
	// 	processar()
	// 	fmt.Println()
	// }

	// execução com Go Routines
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go processar2(i, c)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("programa encerrado em ", <-c)
	}

	// time.Sleep(5000 * time.Millisecond)
	// fmt.Println("Programa encerrado")

	// var v int = 12

	// fmt.Println("antes = ", v)

	// aumentar(&v)

	// fmt.Println("depois = ", v)

	// exemploOs1()
	// exemploOs2()
	// exemploOs3()
	// ponteirosVar()
}
