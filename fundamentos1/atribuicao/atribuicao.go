package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferecnia de tipo
	i += 3
	i -= 3
	i *= 3
	i /= 3
	i %= 2
	fmt.Println(i)

	x, y := 5, 10
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}
