package main

import (
	"fmt"
	"log"

	"github.com/eneassena/bootcampGo/modulo2/200522/exercicio"
)

func main() {
	media, erro := exercicio.MediaAluno(7, 7, -7)
	if erro != nil {
		log.Println(erro.Error())
	}
	fmt.Println("media do aluno é: ", media)
}
