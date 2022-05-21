package main

import "github.com/eneassena/bootcampGo/200522/exercicio"

func main() {

	exercicio.TestStruct()

	// l := pratica.ListaHeterogenea{}
	// l.Data = append(l.Data, 1)

	// l.Data = append(l.Data, true)
	// l.Data = append(l.Data, false)
	// l.Data = append(l.Data, "Eneas")
	// l.Data = append(l.Data, 26)

	// fmt.Println(l.Data)

	// fmt.Println(pratica.Info)
	// pratica.Info = 25
	// fmt.Println(pratica.Info)

	// var i interface{} = "hello"
	// fmt.Println(i)

	// s, ok := i.(string)
	// fmt.Println(s, ok)

	// f, ok := i.(float64)
	// fmt.Println(f, ok)

	// f = i.(float64)
	// fmt.Println(f)

	// f = pratica.Info.(float64)
	// fmt.Println(f)

	// fmt.Println("detalhes do novo triangulo")
	// t := pratica.FiguraGeometrica("triangulo", 2, 5)
	// detalhes(t)

	// fmt.Println("detalhes do novo circulo")
	// t = pratica.FiguraGeometrica("circulo", 5)
	// detalhes(t)

	// fmt.Println("detalhes do triangulo")
	// t := pratica.Triangulo{Largura: 5, Altura: 5}
	// detalhes(t)

	// fmt.Println("detalhes do circulo")
	// c := pratica.Circulo{Raio: 5}
	// detalhes(c)

	// detalhes(c)
	// detahles(c)
	// automovel := pratica.Automovel{}
	// automovel.Correr(360)
	// automovel.Detalhes()

	// moto := pratica.Moto{}
	// moto.Correr(360)
	// moto.Detalhes()

	// exercicio da tarde numero 2
	// exercicio.TestStruct()

	// exercicio da tarde numero 1
	// a := exercicio.Aluno{
	// 	Nome:           "Jonas",
	// 	Sobrenome:      "Sena",
	// 	RG:             "12365498789",
	// 	DataDeAdmissao: "21.05.2022"}

	// exercicio.Detalhes(a)

	// Exercicio 4
	// minimo := exercicio.ExecFunctions("Menor")
	// maximo := exercicio.ExecFunctions("Maior")
	// media := exercicio.ExecFunctions("Media")

	// fmt.Println("Menor dos valor da entra: [10, 4, 5, 6, 7, 8, 9, 56, 4, 34, -45, 2, 3]")
	// fmt.Println(minimo(10, 4, 5, 6, 7, 8, 9, 56, 4, 34, -45, 2, 3))

	// fmt.Println("\nMaior dos valor da entra: [10, 4, 5, 6, 7, 8, 9, 56, 4, 34, -45, 2, 3]")
	// fmt.Println(maximo(10, 4, 5, 6, 7, 8, 9, 56, 4, 34, -45, 2, 3))

	// fmt.Println("\nMédia dos valor da entra: [3, 3, 3]")
	// fmt.Println(media(3, 3, 3))

	// Exercicio 5
	// var amount int

	// animalGato, msgGato := exercicio.Animal("chani") // retorna uma mensagem de erro
	// animalGato, msgGato := exercicio.Animal(GATO) // retorna a função do animal Gato corretamente
	// animalCaes, msgCaes := exercicio.Animal(CAES)
	// animalHamster, msgHamster := exercicio.Animal(HAMSTER)
	// animalTarantula, msgTarantula := exercicio.Animal(TARANTULA)

	// if animalGato != nil {
	// 	amount += animalGato(5)
	// }
	// amount += animalCaes(3)
	// amount += animalHamster(8)
	// amount += animalTarantula(4)

	// if msgCaes != nil || msgGato != nil || msgHamster != nil || msgTarantula != nil {
	// 	fmt.Println("Mensagem pro animal Gato", msgGato)
	// 	fmt.Println("Mensagem pro animal Gato", msgCaes)
	// 	fmt.Println("Mensagem pro animal Hamster", msgHamster)
	// 	fmt.Println("Mensagem pro animal Tarantula", msgTarantula)
	// }

	// fmt.Println("o total de alimento será de = ", amount)

}
