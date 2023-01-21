package exemplos

import (
	"io/ioutil"
)

func NewArquivo() {
	// file, err := os.Create("p1.go")
	// texto := []byte("informação nova")
	// file := ioutil.WriteFile("log/testo.txt", texto, 0644)

	// fmt.Println(file, string(texto))

	// arquivo, err := os.OpenFile("log/arquivo1.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	data := []byte("Manipulando arquivos com Go")
	if err := ioutil.WriteFile("log/texto.txt", data, 0644); err != nil {
		panic(err)
	}
}
