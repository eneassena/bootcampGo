package pratica

import (
	"fmt"
)

/*
	entenda o pacote os
	serve para executar recusos do sistema
*/

func verTodosDir() {
	defer func() {
		fmt.Println("função anonima")
	}()

	d1 := []byte("Hello world")
	// panic("fim da execução")
	fmt.Printf("%c", d1)

}

// func escritaDeArquivo() {
// 	d1 := []byte("ola povo do brasil")

// 	file := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 600)

// }
func PlayPackOs() {
	verTodosDir()
}
