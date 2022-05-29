package tarefas

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
func novoProduto(nome string, preco int, quantidade int) Produto {
	return Produto{Nome: nome, Preco: preco, Quantidade: quantidade}
}

// Função de adicionar produto
func addProduto(user Usuarios, produto Produto) {
	user.produtos = append(user.produtos, produto)
}

// função deletar produto
func deleteProduto(user Usuarios) {
	if len(user.produtos) == 0 {
		fmt.Println("Nao foi possive remover os produtos")
	}
	fmt.Println("Produto removido com sucesso")
	user.produtos = []Produto{}
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

func (u *Usuarios) show() string {
	return fmt.Sprintf("Nome: %s\tSobrenome: %s\tIdade: %v\tEmail: %s\tSenha: %s\n",
		u.Nome, u.Sobrenome, u.Idade, u.Email, u.Senha)
}

func Play() {
	carlos := Usuarios{}
	carlos.setFullName("Carlos", "Oliveira")
	carlos.setEmail("carlosoliveira321@hotmail.com")
	carlos.setIdade(45)
	carlos.setSenha("SDFS_hj465413")
	fmt.Println(carlos.show())

	marcelo := Usuarios{
		Nome:      "Marcelo",
		Sobrenome: "Santos",
		Idade:     35,
		Email:     "marcelo@mail.com",
		Senha:     "123456",
	}
	marcelo.show()

	prod := novoProduto("Caderno", 65, 5)
	addProduto(carlos, prod)
	prod = novoProduto("Tablet", 500, 85)
	addProduto(carlos, prod)

	fmt.Println("Removendo produtos do carlos")
	deleteProduto(carlos)
	fmt.Println(carlos.produtos)

	fmt.Println("Removendo produtos do marcelo")
	deleteProduto(marcelo)
	fmt.Println(marcelo.produtos)
}
