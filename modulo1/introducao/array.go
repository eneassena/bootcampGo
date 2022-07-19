package introducao

import "fmt"

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
