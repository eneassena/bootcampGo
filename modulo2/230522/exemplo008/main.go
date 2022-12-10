package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	// abri um arquivo no diretorio atual
	file, err := os.Open("produto.csv")
	if err != nil {
		log.Fatal(err)
	}
	// antes de finalizar fecha o arquivo
	defer file.Close()

	// ?
	w := tabwriter.NewWriter(os.Stdout, 20, 30, 1, '\t', tabwriter.AlignRight)

	// ?
	scanner := bufio.NewScanner(file)

	// ?
	cabecalho := strings.Split(scanner.Text(), ",")

	for _, c := range cabecalho {
		fmt.Fprintf(w, "%s\t", c)
	}

	fmt.Fprintln(w)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for _, v := range values {
			fmt.Fprintf(w, "%s\t", v)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}
