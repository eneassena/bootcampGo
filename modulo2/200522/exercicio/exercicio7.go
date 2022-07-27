package exercicio

import "fmt"

/* Descrição da atividade 2 da tarde

Exercício 2 - Produtos de e-commerce
Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar
produtos e devolver o valor do preço total.
As empresas têm 3 tipo-de-dados de produtos:
- Pequeno, Médio e Grande.
Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

Lista de custos adicionais:
- Pequeno: O custo do produto (sem custo adicional)
- Médio: O custo do produto 3+ % pela disponibilidade no estoque
- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
adicional pelo envio de $2500.


Requisitos:
- Criar uma estrutura “loja” que guarde uma lista de produtos. [ok]
- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço [ok]
- Criar uma interface “Produto” que possua o método “CalcularCusto” [ok]
- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”. [ok]
- Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome
e preço, e devolva um Produto. [ok]

- Será necessário uma função “novaLoja” que retorne um Ecommerce.
- Interface Produto:
- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
custo adicional segundo o tipo de produto. [ok]

- Interface Ecommerce:
- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
base no custo total dos produtos + o adicional citado anteriormente (caso a categoria
tenha)[ok]

- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
e adicioná-lo a lista da loja [ok]



*/

type Produto struct {
	tipo  string
	nome  string
	preco float64
}

type Loja struct {
	estoque []Produto
}

type ProdutoInterface interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Total() float64
	Adicionar(produto Produto)
}

func (p *Produto) CalcularCusto() float64 {
	switch p.tipo {
	case "Medio":
		acescimo := (p.preco * 0.03)
		p.preco += acescimo
	case "Grande":
		acescimo := (p.preco * 0.06)
		p.preco += acescimo + 2500
	}
	return p.preco
}

func (loja *Loja) Total() float64 {
	var sum = 0.0
	for _, v := range loja.estoque {
		sum += float64(v.CalcularCusto())
	}
	return sum
}

func (loja *Loja) Adicionar(produto Produto) {
	loja.estoque = append(loja.estoque, produto)
}

func novoProduto(novoTipo string, novoNome string, novoPreco float64) Produto {
	return Produto{tipo: novoTipo, nome: novoNome, preco: novoPreco}
}

func novaLoja() Ecommerce {
	return &Loja{}
}

func TestStruct() {
	Mercado := novaLoja()

	produtoPequeno := novoProduto("Pequeno", "carderno", 45.00)
	produtoMedio := novoProduto("Medio", "Livro", 150.00)
	produtoGrande := novoProduto("Grande", "Teclado", 80.00)

	Mercado.Adicionar(produtoPequeno)
	Mercado.Adicionar(produtoMedio)
	Mercado.Adicionar(produtoGrande)

	var tot = Mercado.Total()
	fmt.Println("Total", tot)
	fmt.Println(produtoPequeno.CalcularCusto())
	fmt.Println(produtoMedio.CalcularCusto())
	fmt.Println(produtoGrande.CalcularCusto())
}
