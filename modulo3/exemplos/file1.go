package exemplos

/*
	Objetivo:
	* O que é uma API e para que serve?
	* Protocolos de comunicação(HTTP)
	* Como é composta uma arquitetura web?
	* O que é JSON - Funcionalidade e recursos
	* Nosso primeiro projeto usando Gin
	* Conceitos e características do gin

	- API - é um conjunto de definições e protocolos usados para desenvolver e intergrar
	software de aplicativo. API significa interface de Programação de aplicativos.

	As APIs permitem que seus produtos e serviços se comuniquem com outros, sem a
	cecessidade de saber como eles são implementados

	As APIs (públicas) permitem que as empresas ativem o acesso aos seus recursos,
	mantendo a segurança e o controle

	O que é HTTP?
	"
		HTTP é um protocolo de transferência que permite a comunicação entre o cliente e o
		servidor de forma padronizada.
	"

	Referência disponível em: https://developer.mozilla.org/en-US/docs/Web/HTTP

	Protocolo
	"
		um protocolo pode ser um documento ou um regulamento que estabelece como agir em determinados procedimentos.
	"

	Arquitetura WEB:
		- Cliente-Servidor
		- Interface Uniforme
		- Sistema de camadas
		- Cache
		- Sem estado(Sateless)
		- Código sob demanda

	Sistema de camadas
		O sistema de camadas de chave permite que intermediários na rede, como proxies
		e gateways, sejam implementados de forma tranparente entre um cliente e um servidor
		usando a interface unifo rme da web.

		Os usos mais comuns são para aplicação de segurança, armazenamento em cache
		de respota e balanceamento de carga.

		Cliente 	Verificação do Usuário 		Servidor


	Cache:
		O cache da web é um recurso usado por navegadores e servidores
		para reduzir a quantidade de dados usados ao carregar uma página
		da web

		O cache de dados é uma das chaves mais importantes para a arquitetura da
		web. Ele ajuda a reduzir a latência percebida pelo cliente, aumentar a
		disponibilidade e a confiabilidade de um aplicativo e controlar a carga em um
		servidor web.

		Um cache pode existir em qualquer lugar entre o cliente e o servidor.


	Sem estado (stateless)
		A chave Stateless propõe que todas as interações entre o cliente e o servidor sejam
		tratadas como novas e absolutamente independentes sem salvar o estado.

		Portanto, se quisermos --por exemplo-- que o servidor distinga entre usuários
		logados ou convidados, devemos enviar todas as informações de autenticação necessárias
		em cada solicitação que fizemos a esse servidor.

		Características são:
			* Servidores diferentes fornecem informações diferentes ao mesmo tempo
			* As solicitações anteriores não sõa armazendas
			* As solicitações são independentes, não dependem de resultados anteriores

	Código sob demanda
		O código sob demanda tende a estabelecer um acoplamento tecnológico entre servidores
		web e seus clientes, pois o cliente deve ser capaz de entender e executar o código que baixa
		sob demanda do servidor.

	O que é REST?
	"
		REST (Representational State Tranfer) é o nome que Fielding deu á Composite Web
		Architecture, devido ás chaves vistas acima
	".

	JSON
	"
		Conceito: JSON é a sigla para JavaScriptObjectNotation, ou Notação de Objetos JavaScript. E
		um formato de troca de dados leve.

		é um formato de troca de dados conveniente devido a sua facilidade de
		interoperabilidade com outras linguagens de programação e sua legibilidade
		para os seres humanos
	"

	Sintaxe JSON
	* Uma coleção de pares nome/valore. Em vários idiomas, isso é conhecido como
		objeto, registro, estrutura, dicionario, tabela de hash, lista de chaves ou matriz
		associativa
	* Uma lista ordenada de valores. Na maioria das linguagens, isso é implementado como arrays.
		vetores ou listas

	JSON - pode representar quatro tipo-de-dados primitivos (strings, números, booleanos e valores nulos)
	e dois tipo estruturados (objetos e arrays).

	pacotes JSON


	"
		O pacote JSON é uma lib que nos permite transformar estruturas e tipo-de-dados de dados em GO
		para JSON e vice-versa
	"

	.Marsjal()
	A função func Marshal(v interface{}) ([]byte, error) toma como parâmentro um valor de
	qualquer tipo e retorna uma fatia de butes que contém sua representação no formato JSON.
	Ele também retorna um erro se um for encontrado.


	.Unmarshal()
	esta função  recebe como primeiro parâmentro um array de bytes e como segundo parâmentro
	um ponteiro para uma struct. Se o array de bytes for de dados JSON, a função
	unmarshal() tentará decodificá-la e preencher a estrutura com esses dados.
	A função retorna um erro, se for encontrado.


	Pacotes WEB GO
	net/http
	Este pacote permite gerar servidores web de forma simples.
	um conteiro fundamental em servidores net/http são os handlers. Um manipulador é um ovjeto que implementa a interface
	http.Handler. Uma maneira comun de escrevar um manipulador é usar o adaptador http.HandlerFunc em funções com assinatura
	adequada.

	As funções que servem como manipuladores aceitam um http.ResponseWriter e um http.Request como argumentos. o
	ResponseWriter é usado para retornar a resposta HTTP.

	framework Go = Gin
	é um framework de alto desempenho para criação de aplicações web e microserviços em Go



*/
