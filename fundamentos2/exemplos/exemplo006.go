package exemplos

import (
	"io/fs"
	"io/ioutil"
)

// criando arquivos
func criandoArquivosWithIoUtil(filename string, content []byte, permision fs.FileMode) error {
	err := ioutil.WriteFile(filename, content, permision)
	return err
}

// outra forma de criar aquivo
func criandoArquivosWithOs(filename string, content []byte, permision fs.FileMode) error {
	err := ioutil.WriteFile(filename, content, permision)
	return err
}

type person struct {
	nome     string
	idade    int
	email    string
	telefone string
}

func criatePerson(nome string, idade int, email string, telefone string) person {
	p := person{
		nome:     nome,
		idade:    idade,
		email:    email,
		telefone: telefone,
	}
	return p
}
