package tarefas

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	min = 1
	max = 10000
)

type Cliente struct {
	Arquivo          int    `json:"arquivo"`
	Nome             string `json:"nome"`
	Sobrenome        string `json:"sobrenome"`
	RG               int    `json:"rg"`
	NumeroDeTelefone int    `json:"numero_de_telefone"`
	Endereco         string `json:"endereco"`
}

func (c Cliente) campos() string {
	dados := fmt.Sprintf("\n%d,%s,%s,%d,%d,%s",
		c.Arquivo,
		c.Nome,
		c.Sobrenome,
		c.RG,
		c.NumeroDeTelefone,
		c.Endereco)
	return dados
}

func Error() {
	message := recover()
	texto := fmt.Sprintf("%s", message)
	log.Println(texto)
}

func header() string {
	return "Arquivo,Nome,Sobrenome,RG,NumeroDeTelefone,Endereco"
}

func leiaArquivo(caminho string) *os.File {
	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		panic(fmt.Errorf("error: o arquivo não foi encontrado ou está danificado %w", err))
	}
	return file
}

func geraID() int {
	rand.Seed(time.Now().UTC().UnixNano())
	numero := min + rand.Intn(max-min)
	return numero
}

func registerClientes(file *os.File, cliente []Cliente) error {
	if len(cliente) == 0 {
		defer Error()

		panic(fmt.Errorf("error: o arquivo não foi encontrado ou está danificado"))
	}

	defer file.Close()

	if _, err := file.WriteString(header()); err != nil {
		fmt.Println(err)
		return err
	}

	// add campos de dados do clientes
	for _, c := range cliente {
		if _, err := file.WriteString(c.campos()); err != nil {
			fmt.Println("Erro: registro dos dados do cliente falho")
			return err
		}
	}
	return nil
}

func ExecuteTarefa() {
	clientes := []Cliente{}
	jose := Cliente{
		Arquivo:          geraID(),
		Nome:             "Jose",
		Sobrenome:        "Silva",
		RG:               131654321,
		NumeroDeTelefone: 1321413465,
		Endereco:         "Rua A, Travessa, Numero 520",
	}
	clientes = append(clientes, jose)
	clientes = append(clientes, Cliente{
		Arquivo:          geraID(),
		Nome:             "Carol",
		Sobrenome:        "Silvana",
		RG:               131654321,
		NumeroDeTelefone: 1321413465,
		Endereco:         "Rua A, casa, Numero 520",
	})
	clientes = append(clientes, Cliente{
		Arquivo:          geraID(),
		Nome:             "Carlos",
		Sobrenome:        "Silva",
		RG:               12634625344,
		NumeroDeTelefone: 2346542344,
		Endereco:         "Rua C, Apartamento Bloco D, Porto segunro",
	})

	if err := registerClientes(leiaArquivo("customers.txt"), clientes); err != nil {
		log.Println("Falha no registro dos clientes")
	}

	fmt.Println("Fim da execução")
}

/*func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}*/
