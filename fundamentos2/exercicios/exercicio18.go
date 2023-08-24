package exercicios

import (
	"fmt"
)

type Produtos struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Servicos struct {
	Nome             string
	Preco            float64
	MinutosTrabalhos int
}

type Manutencao struct {
	Nome  string
	Preco float64
}

func SomarProduto(produto []Produtos, c chan int) (totalProduto int, err error) {
	var total = 0

	if len(produto) == 0 {
		return total, fmt.Errorf("impossivel calcular valor dos produto com lista vazia")
	}

	for _, p := range produto {
		total += int(p.Preco * float64(p.Quantidade))
	}
	c <- total
	return total, nil
}

func SomarServico(servicos []Servicos, c chan int) (int, error) {
	somatorio := 0

	if len(servicos) == 0 {
		return 0, fmt.Errorf(fmt.Sprintf("Falha ao ler lista de servicos pois possui %d serviços registrados",
			len(servicos)))
	}

	for _, s := range servicos {
		if s.MinutosTrabalhos < 30 {
			s.MinutosTrabalhos = 30
		}
		somatorio += int(s.Preco * float64(s.MinutosTrabalhos))
	}
	c <- somatorio
	return somatorio, nil
}

func SomarManutencao(manutencao []Manutencao, c chan int) (int, error) {
	total := 0

	if len(manutencao) == total {
		messageErr := fmt.Errorf("error: lista de manutencao esta vazia")
		return total, messageErr
	}

	for _, m := range manutencao {
		total += int(m.Preco)
	}
	c <- total
	return total, nil

}

func ExecuteEmpresaNacional() {

	canalEmpresa := make(chan int)
	orcamentoMensal := 0

	produtos := []Produtos{}
	produtos = append(produtos, Produtos{
		Nome:       "Caneta",
		Preco:      1.99,
		Quantidade: 5,
	})
	produtos = append(produtos, Produtos{
		Nome:       "Caderno",
		Preco:      12.99,
		Quantidade: 3,
	})
	produtos = append(produtos, Produtos{
		Nome:       "Muchila",
		Preco:      50.00,
		Quantidade: 8,
	})
	produtos = append(produtos, Produtos{
		Nome:       "Apontador",
		Preco:      5.99,
		Quantidade: 10,
	})
	produtos = append(produtos, Produtos{
		Nome:       "Livro",
		Preco:      15.59,
		Quantidade: 2,
	})
	produtos = append(produtos, Produtos{
		Nome:       "Agenda",
		Preco:      34.59,
		Quantidade: 1,
	})

	servicos := []Servicos{}
	servicos = append(servicos, Servicos{
		Nome:             "Ajuste Equipamento",
		Preco:            13,
		MinutosTrabalhos: 60,
	})
	servicos = append(servicos, Servicos{
		Nome:             "Support Instalação de Novas Ferramentas",
		Preco:            8,
		MinutosTrabalhos: 35,
	})
	servicos = append(servicos, Servicos{
		Nome:             "instalação de novos equipamento de ambiente",
		Preco:            20,
		MinutosTrabalhos: 120,
	})
	servicos = append(servicos, Servicos{
		Nome:             "desenvolvimento de novas funções",
		Preco:            5,
		MinutosTrabalhos: 240,
	})

	manutencao := []Manutencao{}
	manutencao = append(manutencao, Manutencao{
		Nome:  "Manutenção no software",
		Preco: 950,
	})
	manutencao = append(manutencao, Manutencao{
		Nome:  "Ativa Segurança ofensiva",
		Preco: 800.65,
	})
	manutencao = append(manutencao, Manutencao{
		Nome:  "Gestão de travego na rede",
		Preco: 300.02,
	})
	manutencao = append(manutencao, Manutencao{
		Nome:  "cadastro de novo gesto de bakup",
		Preco: 100,
	})

	go SomarProduto(produtos, canalEmpresa)
	orcamentoMensal += <-canalEmpresa

	go SomarServico(servicos, canalEmpresa)
	orcamentoMensal += <-canalEmpresa

	go SomarManutencao(manutencao, canalEmpresa)
	orcamentoMensal += <-canalEmpresa

	fmt.Printf("gasto total mensal %.2f\n", float64(orcamentoMensal))
}
