package mapping

import (
	"fmt"
	"sort"
)

func Mapeamento3() {
	ages := map[string]int{
		"Raquel": 32,
		"Sheila": 43,
		"Carlos": 54,
		"Ana":    23,
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s - %v\n", name, ages[name])
	}
}
