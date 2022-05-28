package exercicio

import "fmt"

/* Descrição da atividade 1 da tarde

 Exercício 1 - Registro de estudantes

Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os
detalhes dos dados de cada um deles, conforme o exemplo abaixo:
Nome: [Nome do aluno]
Sobrenome: [Sobrenome do aluno]
RG: [RG do aluno]
Data de admissão: [Data de admissão do aluno]

Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos
alunos.
Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG,
Data e que tenha um método de detalhamento



*/

type Aluno struct {
	Nome           string
	Sobrenome      string
	RG             string
	DataDeAdmissao string
}

func Detalhes(aluno Aluno) {
	fmt.Printf("Nome: \t\t\t%s\n", aluno.Nome)
	fmt.Printf("Sobrenome: \t\t%s\n", aluno.Sobrenome)
	fmt.Printf("RG: \t\t\t%s\n", aluno.RG)
	fmt.Printf("Data de Admissão: \t%s\n", aluno.DataDeAdmissao)
}
