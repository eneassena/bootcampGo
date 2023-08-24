package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a+-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b)

	fmt.Println("AND = ", a&b)
	fmt.Println("OR = ", a|b)
	fmt.Println("XOR = ", a^b)

	c := 3.0
	d := 2.0
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor => ", math.Min(c, d))
	 
}
