// pacote do programa
package introducao

// importe dos pacote externos
import "fmt"

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
}
