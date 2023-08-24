package exercicios

import (
	"fmt"
	"math/rand"
)

func selection_sort(arr_original []int, canal chan []int) {
	var j int
	var i int
	for j = 0; j < len(arr_original); j++ {
		for i = j; i < len(arr_original); i++ {
			if arr_original[j] > arr_original[i] {
				aux := arr_original[i]
				arr_original[i] = arr_original[j]
				arr_original[j] = aux
			}
		}
	}
	canal <- arr_original
	// return arr_original
}

func insertion_sort(arreglo []int, canalOrder chan []int) {
	t := len(arreglo)
	for i := 1; i < t; i++ {
		valor_insert := arreglo[i]
		j := i - 1

		for j >= 0 && arreglo[j] > valor_insert {
			arreglo[j+1] = arreglo[j]
			j--
		}
		arreglo[j+1] = valor_insert
	}
	canalOrder <- arreglo
}

func PlayTarefasWithCanal() {
	var (
		arr_original_100  = rand.Perm(100)
		arr_original_1000 = rand.Perm(1000)
	///	arr_original_10000 = rand.Perm(10000)
	)

	canal := make(chan []int)

	go selection_sort(arr_original_100, canal)
	arr_ordenado := <-canal
	fmt.Println("selection sort = ", arr_ordenado)

	go insertion_sort(arr_original_1000, canal)
	arr_ordenado = <-canal
	fmt.Println("insertion sort = ", arr_ordenado)

	//arr_order_insertion := insertion_sort(arr_original_1000)
	// arr_order_insertion2 := insertion_sort(arr_original_10000)

	//fmt.Println("Ordenação por insertion sort = ", arr_order_insertion)
	//fmt.Println("Ordenação por insertion 2 sort = ", arr_order_insertion2)
	// fmt.Println("Ordenação por selection sort = ", arr_order_selection)
}
