### objetivo 
- entender o pacote fmt
- entender o pacote os
- entender o pacote io

### caracteres de scape 
- \n pula uma linha
- \\ barra invertida
- \t tabulação 
- \v tabulação vertical

### tentar acessar printenv
no mac ver variavel de sistema

### comando log
usamos o log.SetFlags(0) para retira a data que imprimida no .log

### entenda sobre os codigos de permissão no linux
https://chmod-calculator.com/


### defer para a execução no final
panic()
para a execução do sistema

### criando função anonima

defer func() {
	fmt.Println("função anonima")
}()



#### Para definir o valor de uma variável de ambiente, use o comando de shell apropriado para associar o nome da variável a um valor. Por exemplo, para definir a variável PATH no valor /bin:/sbin:/user/bin:/user/sbin:/system/Library/, o seguinte comando deve ser digitado em uma janela do Terminal:
$ PATH=/bin:/sbin:/user/bin:/user/sbin:/system/Library/ export PATH

### Para ver todas as variáveis de ambiente, digite:
% env


### site para referencias
http://goporexemplo.golangbr.org/