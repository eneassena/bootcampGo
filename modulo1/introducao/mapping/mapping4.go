package mapping

import "fmt"

func Mapeamento4() {
	ages := map[string]int{"age1": 1, "age2": 2, "age3": 3}

	names := make([]string, 0, len(ages))

	fmt.Println(names)
}
