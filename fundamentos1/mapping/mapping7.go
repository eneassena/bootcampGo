package mapping

import "fmt"

func Mapeamento7() {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	b := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	err := equal(a, b)
	fmt.Println(err)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
