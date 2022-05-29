package pratica

import (
	"fmt"
	"time"
)

/*
	vamos estudar criação de canal em go
	para definir um canal do tipo inteiro fazemos da seguinte forma:
	usamos a palavra reservada chan
	c := make(chan int)

	um ponteiro é um endereço de memória que referencia outro valor

	O que é simultaneidade e paralelismo?
	Simutaneidade é o ato de iniciar, executar e concluir duas ou mais tarefas em períodos de tempo
	escalonados. Não significa necessariamente que ambos estão sendo executados ao mesmo tempo

	Parelelismo implica que duas ou mais tarefas sejam executdas exatamente ao mesmo tempo.

	Conceito de GO routines?
	Eles são a solução oferecida pela GO para implementar a simutaneidade.
	Não são threads! GO routines são gerenciadas pelo GO Runtime e seu agendador, não pelo sistema operacional
	Quando executamos uma função como GO routine, sua execução não bloqueará a continuação da função que a 
	chamou, pois ela será executada simultaneamente com esta

*/

func processar(i int, c chan int) {
	fmt.Println(i, "Inicia")
	fmt.Println(1000 * time.Millisecond)
	fmt.Println(i, "Termina")
	c <- i
}

func Ok(){
	c := make(chan int)
	for i := 1; i <= 10; i++ {
		go processar(i, c)
	}
	for i := 1; i <= 10; i++ {
		variavel := <-c 
		// obtendo o valor do canal
		fmt.Println("valor do canal: ", variavel)
	}
	
	fmt.Println("Programa encerrado")
}