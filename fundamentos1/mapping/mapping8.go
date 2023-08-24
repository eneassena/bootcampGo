package mapping

import "fmt"

var m = make(map[string]int)

func Mapeamento8() {
	list := []string{"a", "b", "c"}
	res := k(list)
	add(list)
	add(list)
	add(list)
	fmt.Println(Count(list), res)
	fmt.Println("Map = ", m)
}

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
