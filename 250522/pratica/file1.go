package pratica

import (
	"fmt"
	"log"
)

type MyError struct{}

func (e MyError) ErrorHTTP(code int) string {
	switch code {
	case 100:
		return "Error de informacoes"
	case 200:
		return "error de sucesso"
	case 300:
		return "error de permissao"
	case 400:
		return "error de cliente"
	case 500:
		return "erro de servidor"
	}
	return "erro fora do comun"
}

func Play() {
	err := MyError{}

	fmt.Println(err.ErrorHTTP(200))
	log.Println(err.ErrorHTTP(500))
	fmt.Print("codigo executavel Go\n")
}
