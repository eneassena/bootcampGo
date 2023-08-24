package exercicios

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// A denifição de uma estrutura se da pela: determinação de seus campos seguidos por
// um espaço e o tip de dados
// para separa os campos usa-se uma quebra de linha

type Preferencias struct {
	Comidas  string
	Filme    string
	Series   string
	Animes   string
	Esportes string
}

type Pessoa struct {
	Nome         string
	Genero       string
	Idade        int
	Profissao    string
	Peso         float64
	Preferencias Preferencias
}

type ProdutoLoja struct {
	Nome                string  `json:"nome_produto" bd:"NameProduto"`
	Preco               float64 `json:"preco_produto" bd:"PriceProduto"`
	QuantidadeEmEstoque int     `json:"quantidade_em_estoque" bd:"CountInEstoque"`
}

func AnalisandoEstrutra() {

	produto := ProdutoLoja{"Teclado", 50, 80}
	reflectProduto := reflect.TypeOf(produto)

	fmt.Println(produto)
	fmt.Println("Name:", reflectProduto.Name())
	fmt.Println("Kind:", reflectProduto.Kind())
	fmt.Println("Size:", reflectProduto.Size())
	fmt.Println("NumField:", reflectProduto.NumField())

	for i := 0; i < reflectProduto.NumField(); i++ {
		field := reflectProduto.Field(i)
		fmt.Println(field)
		// acessando a tag/rotulo do struct
		tagBD := field.Tag.Get("bd")
		tagJSON := field.Tag.Get("json")
		fmt.Println(tagBD)
		fmt.Println(tagJSON)
	}

}

func ExecuteEstrutura() {
	produto := ProdutoLoja{"Teclado", 60, 50}
	fmt.Println(produto)

	meuJson, err := json.Marshal(produto)

	fmt.Println("ProdutoLoja em Json")

	data := string(meuJson)

	fmt.Println(data, err)

	prod := ProdutoLoja{}
	p := reflect.TypeOf(prod)
	fmt.Println("Name", p.Name())
	fmt.Println("Kind", p.Kind())

	pe := Pessoa{"Eneas", "Masculino", 25, "Programador", 65, Preferencias{}}
	pe.Preferencias = Preferencias{Comidas: "Verduras", Animes: "Cachorro"}
	play(pe)

	fmt.Println()

	pe2 := Pessoa{"Ana", "Feminino", 22, "Dentista", 70, Preferencias{"Lasanha", "Transforme 4", "Os Vincks", "Gatos", "Basquete"}}
	play(pe2)

	fmt.Println("Olá", pe2.Nome, "Feliz aniversário")
	pe2.Idade = 23
	play(pe2)

}

func play(pessoa Pessoa) {
	fmt.Printf("Nome: %s\nGenero: %s\nIdade: %v\nProfissão: %s\nPeso: %v\n\n%s\nComida: %s\nFilme: %s\nSeries: %s\nAnimal: %s\nEsporte: %s\n",
		pessoa.Nome, pessoa.Genero, pessoa.Idade, pessoa.Profissao, pessoa.Peso, "Preferencias",
		pessoa.Preferencias.Comidas,
		pessoa.Preferencias.Filme,
		pessoa.Preferencias.Series,
		pessoa.Preferencias.Animes,
		pessoa.Preferencias.Esportes,
	)
}
