package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings: ", "LAPIS" == "LAPIS")
	fmt.Println("int: ", 5 < 10)
	fmt.Println("int: ", 5 <= 10)
	fmt.Println("int: ", 5 > 10)
	fmt.Println("int: ", 5 >= 10)
	fmt.Println("int: ", 5 != 10)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println(d1 == d2)
	fmt.Println(d1.Equal(d2))

	type Pessoa struct {
		Name string
	}

	p1 := Pessoa{Name: "Pedro"}
	p2 := Pessoa{Name: "Pedro"}
	fmt.Println(p1 == p2)
	p2.Name = "Ana"
	fmt.Println(p1 != p2)
}
