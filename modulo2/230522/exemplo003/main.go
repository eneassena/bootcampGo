package main

import "fmt"

type Greeting interface {
	Hello() string
}

type Portuguese struct{}

func (p Portuguese) Hello() string {
	return "Olá"
}

type English struct{}

func (p English) Hello() string {
	return "Hello"
}

func SayHello(g Greeting) string {
	return g.Hello()
}

func main() {
	port := Portuguese{}
	eng := English{}

	fmt.Println(SayHello(port))
	fmt.Println(SayHello(eng))
}
