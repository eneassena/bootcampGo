package mapping

import "fmt"

func Mapeamento6() {
	var ages map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(ages)
	age, ok := ages["a"]
	fmt.Println(age, ok)
	if ok {
		fmt.Println(age)
	}
}
