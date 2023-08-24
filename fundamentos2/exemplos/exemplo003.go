package exemplos

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

func SayHello009(g Greeting) string {
	return g.Hello()
}
