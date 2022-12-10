package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Id    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

type Response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// struct para json
	response1 := &Response{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	// converte a Response em um JSON
	byteResponse1, _ := json.Marshal(response1)
	fmt.Println(string(byteResponse1))

	response2 := Response{Page: 1, Fruits: []string{"laranjada", "maçã"}}
	byteResponse2, _ := json.Marshal(response2)
	fmt.Println(string(byteResponse2))

	// cria um product
	product1 := []Product{
		{Id: 1, Nome: "Teclado", Preco: 15.00},
		{Id: 2, Nome: "Sapato", Preco: 19.00},
		{Id: 3, Nome: "Tenis", Preco: 30.00},
	}

	// converte o product em JSON
	byteProduct2, _ := json.Marshal(product1)
	fmt.Println(string(byteProduct2))

	// Json para struct
	var product3 []Product

	if err := json.Unmarshal(byteProduct2, &product3); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Voltando para strcut o json tamanho da lista de produtos: ", len(product3))
	for _, product := range product3 {
		fmt.Printf("\nID: %d\tNome: %s\t%.2f\n", product.Id, product.Nome, product.Preco)
	}

	byteData := []byte(`{"Nome":"Hospagem", "preco": 12.99}`)
	var dataMap = make(map[string]interface{})
	if err := json.Unmarshal(byteData, &dataMap); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataMap)

}
