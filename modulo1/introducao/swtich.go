package introducao

import "fmt"

func SwitchCase(dias int) {
	switch dias {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fervereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abriu")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubo")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Desconhecido")
	}
}
