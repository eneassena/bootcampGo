package main

import "fmt"

type Usuarios struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Senha     string
	produtos  []Produto
}

type Produto struct {
	Nome       string
	Preco      int
	Quantidade int
}

// função de Novo Produtos
func (u *Usuarios) novoProduto(nome string, preco int, quantidade int) Produto {
	return Produto{Nome: nome, Preco: preco, Quantidade: quantidade}
}

// Função de adicionar produto
func (u *Usuarios) addProduto(produto Produto) {
	u.produtos = append(u.produtos, produto)
}

// função deletar produto
func (u *Usuarios) deleteProduto() {
	if len(u.produtos) == 0 {
		fmt.Println("Nao foi possive remover os produtos")
	}
	fmt.Println("Produto removido com sucesso")
	u.produtos = []Produto{}
}

func (u *Usuarios) setFullName(nome string, sobrenome string) {
	u.Nome = nome
	u.Sobrenome = sobrenome
}

func (u *Usuarios) setIdade(idade int) {
	u.Idade = idade
}

func (u *Usuarios) setEmail(email string) {
	u.Email = email
}

func (u *Usuarios) setSenha(senha string) {
	u.Senha = senha
}

func (u *Usuarios) showUsuario() {
	fmt.Printf("\nNome: %s\r\nSobrenome: %s\r\nIdade: %v\r\nEmail: %s\r\nSenha: %s\n",
		u.Nome, u.Sobrenome, u.Idade, u.Email, u.Senha)
}

func (u *Usuarios) showProduto() {
	if len(u.produtos) > 0 {
		fmt.Println("\nDados do Produto:")
		for _, value := range u.produtos {
			fmt.Printf("Nome: %s\tPreco: %v\tQuantidade: %v\n", value.Nome, value.Preco, value.Quantidade)
		}
	} else {
		fmt.Println("Não há produtos para esse usuário")
	}
}

func main() {
	carlos := Usuarios{}
	carlos.setFullName("Carlos", "Oliveira")
	carlos.setEmail("carlosoliveira321@hotmail.com")
	carlos.setIdade(45)
	carlos.setSenha("SDFS_hj465413")
	carlos.showUsuario()

	marcelo := Usuarios{
		Nome:      "Marcelo",
		Sobrenome: "Santos",
		Idade:     35,
		Email:     "marcelo@mail.com",
		Senha:     "123456",
	}
	marcelo.showUsuario()

	carlosProd := carlos.novoProduto("Caderno", 65, 5)
	carlos.addProduto(carlosProd)

	marceloProd := marcelo.novoProduto("Tablet", 500, 85)
	marcelo.addProduto(marceloProd)

	carlos.showProduto()
	marcelo.showProduto()

	fmt.Println("\nRemovendo produtos do carlos")
	carlos.deleteProduto()

	fmt.Println("\nRemovendo produtos do marcelo")
	marcelo.deleteProduto()
}
