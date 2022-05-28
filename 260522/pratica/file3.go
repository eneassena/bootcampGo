package pratica

import (
	"errors"
	"fmt"
)

/*
	Iremos estudar o pacote de erros que possui as funções

	- New()		- cria erros que contêm apenas um texto de mensagem

	- As(err error, target interface{}) bool
		* verifica se um erro é de um tipo específico, com
			retorno boolean caso encontre a correspondencia
			retorna true, caso contrario false

	- func Is(err, target interface{}) bool -
		* Compara um erro com um valor. se encontrar uma correspondência, retorna true
			caso contrário, retorna false

	- Unwrap(err error) error	- Um error que contém outro pode implementar esta função.
		Esta função retornao erro subjacente. Se e1 Unwrap() retorna e2,
		então dizemos que e1 envolve e2 e que e1 pode ser desempacotado para obter e2.

	obs.:
		* as funções mult-retorno do GO permitem retornar, entre outros valores, um erorr.
			Esse erro pode ser um que nós mesmos criamos ou um dos errors que essas funções retornam
			por padrão
		* Diferente de outras linguagens que usam exceções para tratamento de muitos errors, é idiomático em Go
			usar valores de retorno que indicam errors sempre que possível.
		* Ao retornar um erro de uma fnção com vários
			valores de retorno, ocódigo idiomático GO também
			definirá um "valor zero", para cada valor "non-error".
			Os valores zero são por exemplo, uma string vazia
			para string, "0" para inteiros, uma struct vazia para
			tipos de estrutura e "nil" para tipos de ponteiro e
			interface, entre outros

	Boas práticas:
		* Na prática, geralmente não criamos um erro para usar diretamente. Costumamos cria-lo para ser
			usado como um retorno de função
		* Por convenção, quando você usa funções que retornam mais de um valor, o erro deve
			ser retornado por último.

	// criando funções de error
	func FuncName() (string, error) {
		return "", nil



	}
*/

type myError struct {
	msg string
	x   string
}

func (e *myError) Error() string {
	return fmt.Sprintf("ocorreu um error: %s, %s\n", e.msg, e.x)
}

var err1 = errors.New("error número 1")

func x() error {
	return fmt.Errorf("informação extra do código %w", err1) // func Errorf retorna um error
}

type errorTwo struct{}

func (e errorTwo) Error() string {
	return "error two happened"
}

func SayHello(nome string) (string, error) {
	if nome == "" {
		return "", errors.New("nome não poder ser vasio")
	}
	return fmt.Sprintf("Hello: %s", nome), nil
}

func PlayErrorFormat() {
	//nome, err := SayHello("Eneas")
	if nome, err := SayHello("ana"); err != nil {
		fmt.Printf("erorr: %s", err)
	} else {
		fmt.Println(err)
		fmt.Println(nome)
	}

	e2 := errorTwo{}
	e1 := fmt.Errorf("e2: %w", e2)
	e3 := e2
	fmt.Println(errors.Unwrap(e1))
	fmt.Println(errors.Unwrap(e2))
	fmt.Println(errors.Unwrap(e3))

	/*e := x()
	coincidence := errors.Is(e, err1)
	fmt.Println(coincidence)*/

	// e := &myError{msg: "operação invalida", x: "404"}
	// var err *myError

	/*defer func() {
		saidaErr := recover()
		if saidaErr != nil {
			fmt.Println(saidaErr)
		}
	}()*/
	// isMyError := errors.As(e, &err)
	// fmt.Println(isMyError)

	// err := fmt.Errorf("momento do error: %v", time.Now())
	// fmt.Println("error: ", err)
}
