package main

import "fmt"

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2
	sorvete := trabalho1 || trabalho2

	return comprarTv50, comprarTv32, sorvete
}

func main() {
	comptv50, comptv32, sorv := compras(false, false)
	fmt.Printf("Comprar Tv 50: %t, Comprar tv 32: %t, tomar sorvete: %t", comptv50, comptv32, sorv)
}
