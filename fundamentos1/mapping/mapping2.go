package mapping

import "fmt"

func Mapeamento2() {
	// forma 1 de declaração de map
	ages1 := make(map[string]int)
	ages1["carlos"] = 52
	ages1["ana"] = 25
	fmt.Println("carlos tem,", ages1["carlos"], "anos")
	fmt.Println("ana tem,", ages1["ana"], "anos")

	// forma 2 de declaração de map
	ages2 := map[string]int{
		"alice":  34,
		"charle": 34,
	}
	fmt.Println(ages2)

	// criando map vazio
	ages3 := map[string]int{}
	fmt.Println(ages3)

	// removendo uma chave do map
	ages4 := map[string]int{
		"Carol": 19,
		"Pedro": 25,
		"Luiza": 35,
	}
	fmt.Println("lista de idades")
	fmt.Println(ages4)
	nome := "Carol"
	fmt.Println("Removendo Carol da lista")
	delete(ages4, nome)
	fmt.Println(ages4)

	ages4["bob"] = ages4["bob"] + 1
	fmt.Println(ages4)
	ages4["bob"]++
	fmt.Println(ages4)

	//_ = &ages4["bob"] error não é possivel obter o endereço de memoria

	for name, age := range ages4 {
		fmt.Println("Chave:", name, "age:", age)
	}
}
