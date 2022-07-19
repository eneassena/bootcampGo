// pacote do programa
package introducao

// importe dos pacote externos
import (
	"fmt"

	"github.com/eneassena/bootcampGo/modulo2/190522/exercicios"
)

var idade int = 25

// criando contante
const pi float32 = 3.1415

const (
	idade1 = 15
	idade2 = 35
	idade3 = 14
)

// função principal do programa
func Introducao() {
	sobrenome := "Sena"

	// fmt.Println("Hello World")
	var nome string

	nome = "eneas"

	fmt.Println(nome, sobrenome)

	fmt.Println(nome, idade, pi)

	fmt.Println("\nconstantes")
	fmt.Println(idade1, idade2, idade3)

	exercicios.Exercicio8()

	exercicios.Exercicio1()
	exercicios.Exercicio2()
	exercicios.Exercicio3()
	exercicios.Exercicio4()

	// exemplo de ponteiro em Go
	// var idade int = 10
	// var p *int

	// p = &idade

	// fmt.Println(&p) // retorna o endereco me memoria da variavel
	// fmt.Println(*p) // retorna o valor da variavel na momeria
}
