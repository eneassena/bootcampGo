package mapping

import "fmt"

func Mapeamento5() {
	var ages map[string]int
	fmt.Println(ages)
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)

	ages["carol"] = 21
}
