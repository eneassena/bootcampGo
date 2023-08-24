package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type product struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Published bool   `json:"published"`
}

func main() {

	p := []product{
		{
			Name:      "MackBook Pro",
			Price:     1500,
			Published: true,
		},
		{
			Name:      "MackBook Pro 2",
			Price:     1900,
			Published: true,
		},
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("imprime o array de produto em formato json")
	fmt.Println(string(jsonData))

	var newProduto []product

	// converte um json para estruct go
	err = json.Unmarshal(jsonData, &newProduto)
	if err != nil {
		log.Fatal(err)
	}

	// le os dados da lista de produtos
	for _, p := range newProduto {
		fmt.Printf("Name: %s\tPrice: %.2f\tPublished: %t\n", p.Name, float64(p.Price), p.Published)

	}
}
