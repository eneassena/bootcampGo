
Entrada de dados do telcado
fmt.Scan(&variavel)

pesquisar mais sobre ellipsis
essa tecnica é semelhante ao * no python
*args recebe uma tupla de paramentros
exemplo args=(4,5,26,8,56,6)


para tratar erros usamos o pacote do go chamado errors é necessário importa o pacote


codigo em golang ferramenta criada pelo professor
https://github.com/cyruzin/trailer

responder pesquisa digitalhouse

conceito de função 
"A função é um pedaço de código que executa uma tarefa específica.
Todos os programas têm pelo menos uma função principal. As função são diferenciadas umas das outras por meio de indentificador que as torna únicas"


dando continuidade na leitura do mateiral parte da tarde


objetivos:
* Conhecça e aplique as estruturas em Go.
* Compreender e usar métodos dentro de nossas estruturas
* Compreender e aplicar os rótulos das estruturas 
* Conheça e use interface em Go

o que define uma estrutura são é a coleção de campos de dados contido na mesma.


Conceito de Rótulo
Dentro de nossas estruturas podemos definir rótulos ou anotações que se referem a cada um dos cmapos que aparecem após a declaração do tipo de dados


para acessar os rótulo das estruturas usamos o reflect, este pacote nos fornece funcionalidades para obter informações dos objetos em tempo de execução.

Com o método NumField podemos obter o número de cmapos que temos nossa estrura

com o método Field podemos obter o campo da nossa estrutura passando o índice como parâmentro


Go não possui classes. No entanto, os métodos podem ser definidos em tipos de dados.
Um método é uma função com um argumento de receptor especial. O receptor aparece em sua própria lista de argumentos entre a palavra-chave func e o nome do método.



