package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")

	fmt.Println("Pula uma ")
	fmt.Println(" linha")

	y := 500

	fmt.Println("O valor de y é ", y)

	texto := fmt.Sprint("a data atual é = ", "15-05-2022")
	fmt.Println(texto)

	texto = fmt.Sprintf("Esta é uma forma de escrita %s", "formatada")
	fmt.Println(texto)

	fmt.Println("O ano atual é = ", 2022)

	fmt.Printf("int %d \t float64 %f \t float64 com 2 %.2f \t %t \t %s", 10, 10.2122, 20.10000, true, "Golang Google")

	texto = fmt.Sprintf("\nint %d \t float64 %f \t float64 com 2 %.2f \t %t \t %s", 10, 10.2122, 20.10000, true, "Golang Google")
	fmt.Println(texto)

	// imprimindo o tipo das variaveis
	var (
		numeroInt     = 10
		numeroFloat64 = 423.3434233
		message       = "O rato roeu a roupa do rei de roma"
		status        = true
	)
	messageFormat := fmt.Sprintf("\nInt = %d type = %T\nFloat64 = %f type = %T\nString = %s type = %T\nbool = %t type = %T",
		numeroInt, numeroInt,
		numeroFloat64, numeroFloat64,
		message, message,
		status, status)
	fmt.Println(messageFormat)
}
