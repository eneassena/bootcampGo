// pacote do programa
package introducao

// importe dos pacote externos
import (
	"fmt"

	"github.com/eneassena/bootcampGo/190522/exercicios"
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

func ArraySintaxe() {
	var nomes [2]string
	idades := make([]int, 2)

	idades[0] = 15
	idades[1] = 56

	idades = append(idades, 27)

	fmt.Println(idades[0])
	fmt.Println(idades[1])
	fmt.Println(idades[2])

	nomes[0] = "Eneas"
	nomes[1] = "senha"

	fmt.Println("valor dos array")
	fmt.Println(nomes[0])
	fmt.Println(nomes[1])

	fmt.Printf("tamanho do array de idades %v")
	fmt.Printf("tamanho do array de nomes")
}

func Mapeamento() {
	// var studentes = map[string]int{"joao": 20, "Pedro": 26}
	// fmt.Println(studentes["joao"])
	// studentes["Ana"] = 35
	// studentes["Maria"] = 42
	// fmt.Println(studentes["joao"])

	alunos := make(map[string]interface{})
	alunos["Nome"] = "Pedro"
	alunos["Idade"] = 30
	fmt.Println(alunos)

	delete(alunos, "Idade")
	fmt.Println(alunos)
}

func EstruturaRepeticao() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	// frutas := []string{"Maçã", "Melão", "Laranja"}

	// for key, value := range frutas {
	// 	fmt.Println(key, value)
	// }

	// sum := 0
	// for {
	// 	sum += 1
	// 	fmt.Printf("%v\n", sum)
	// }

	// sum := 1
	// for sum < 100 {
	// 	fmt.Println(sum)
	// 	sum += 10
	// }
	// var sum int8 = 0
	// for {
	// 	fmt.Println(sum)
	// 	sum++
	// }

	sum := 0
	for {
		sum++

		fmt.Println(sum)

		if sum%2 == 0 {
			continue
		}
		if sum > 100 {
			break
		}
	}

}

func switchCase() {
	dias := 1
	switch dias {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fervereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abriu")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubo")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Desconhecido")
	}
}
