package pratica

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// A denifição de uma estrutura se da pela: determinação de seus campos seguidos por um espaço e o tip de dados
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

type Produto struct {
	Nome                string  `json:"nome_produto" bd:"NameProduto"`
	Preco               float64 `json:"preco_produto" bd:"PriceProduto"`
	quantidadeEmEstoque int     `json:"quantidade_em_estoque" bd:"CountInEstoque"`
}

func AnalisandoEstrutra() {

	produto := Produto{"Teclado", 50, 80}
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
	produto := Produto{"Teclado", 60, 50}
	fmt.Println(produto)

	meuJson, err := json.Marshal(produto)

	fmt.Println("Produto em Json")

	data := string(meuJson)

	fmt.Println(data, err)

	prod := Produto{}
	p := reflect.TypeOf(prod)
	fmt.Println("Name", p.Name())
	fmt.Println("Kind", p.Kind())

	// p := Pessoa{"Eneas", "Masculino", 25, "Programador", 65, Preferencias{}}
	// p.Preferencias = Preferencias{Comidas: "Verduras", Animes: "Cachorro"}
	// play(p)

	// fmt.Println()

	// p2 := Pessoa{"Ana", "Feminino", 22, "Dentista", 70, Preferencias{"Lasanha", "Transforme 4", "Os Vincks", "Gatos", "Basquete"}}
	// play(p2)

	// fmt.Println("Olá", p2.Nome, "Feliz aniversário")
	// p2.Idade = 23
	// play(p2)

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
