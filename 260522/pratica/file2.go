package pratica

import (
	"fmt"
	"log"
	"os"
)

type myCustomError struct {
	status  int
	message string
}

func (m *myCustomError) Error() string {
	return fmt.Sprintf("Error: %d\tMessage: %s\n", m.status, m.message)
}

func myCustomErrorTest(status int) (int, error) {
	if status >= 400 {
		return 500, &myCustomError{
			status:  status,
			message: "Algo deu errado",
		}
	}
	return 200, nil
}

func Play() {

	status, err := myCustomErrorTest(500)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println(status)
}

