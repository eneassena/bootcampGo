package main

import (
	"io"
	"log"
	"os"
)

func main() {
	_, err := io.WriteString(os.Stdout, "Hello World\n")

	if err != nil {
		log.Fatal(err)
	}
}
