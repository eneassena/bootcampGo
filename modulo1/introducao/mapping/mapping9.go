package mapping

import "fmt"

var graph = make(map[string]map[string]bool)

func Mapeamento9() {
	from := "jose@gmail.com"
	to := "maria@hotmail.com"

	addEdge(from, to)
	fmt.Println(hasEdge(from, to))
}

func addEdge(from, to string) {
	adges := graph[from]
	if adges == nil {
		adges = make(map[string]bool)
		graph[from] = adges
	}
	adges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
