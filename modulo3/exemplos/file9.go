package exemplos

import "os"

// lendo argumentos na linha de comando
func ReadArgs() string {
	var s, sep string 
	for i :=0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}