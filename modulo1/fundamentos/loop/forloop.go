package loop

import "fmt"

func EstruturaRepeticao() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	// frutas := []string{"Maçã", "Melão", "Laranja"}

	// for key, value := range frutas {
	// 	fmt.Println(key, value)
	// }

	// sum := 0
	// for {
	// 	sum += 1
	// 	fmt.Printf("%v\n", sum)
	// }

	// sum := 1
	// for sum < 100 {
	// 	fmt.Println(sum)
	// 	sum += 10
	// }
	// var sum int8 = 0
	// for {
	// 	fmt.Println(sum)
	// 	sum++
	// }

	sum := 0
	for {
		sum++

		fmt.Println(sum)

		if sum%2 == 0 {
			continue
		}
		if sum > 100 {
			break
		}
	}
}
