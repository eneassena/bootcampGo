package main

import (
	"fmt"
	"os"
)

// lendo argumentos na linha de comando
func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Printf("Foram: %d argumentos line passados\n", len(os.Args))
	fmt.Println(s)
}
