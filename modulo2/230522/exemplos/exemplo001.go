package exemplos

import "fmt"

func verboString() {
	fmt.Printf("\nOlá %s\nSeu caractere é = %c", "Eneas", 'A')
	fmt.Print("\nOlá ", "Eneas ", "Seu caractere é = ", string('A'))
}

func saidaWithSprintf() {
	txt := fmt.Sprintf("\nOla %s, seu RG = %v\n", "Eneas", 12345678998)
	fmt.Println(txt)
}

func showTypeof() string {
	return fmt.Sprintf("int: %T\nstring: %T\nfloat64: %T\nBool: %T\n", 5, "ola", 5.10, false)
}
func exemplo1() {
	typeVars := showTypeof()
	fmt.Println(typeVars)

	fmt.Printf("%t\n", true)

	ok, err := fmt.Print(1, 2, 3)
	if err != nil {
		fmt.Println("Error gerado: ", err)
	}
	fmt.Println("\n", ok)

	// fmt.Printf("%10.2f", 1222.13215)
	fmt.Printf("%10d", 132)
	fmt.Printf("%10s", "asd")
}

func exemplo2() {
	nome, age := "kin", 22
	res := fmt.Sprint(nome, " tem, ", age, " anos de idade\n")
	fmt.Println(res)
}

func exemplo3() {
	nome, age := "kin", 22
	res := fmt.Sprintf("Nome: %s, tem %d anos de idade\n", nome, age)
	fmt.Println(res)
}

func Execute() {
	exemplo3()

	exemplo2()
}
