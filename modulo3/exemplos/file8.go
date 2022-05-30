package exemplos


import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
}

type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func PlayExemploStructForJson() { 

	// struct para json
	res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"},
	}
    
	res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

	prod1D := response2{Page: 1, Fruits: []string{"laranjada", "maçã"}}
	prod2B, _ := json.Marshal(prod1D)
	fmt.Println(string(prod2B))

	p1 := []Product{
		{Id: 1, Nome: "Teclado", Preco: 15.00},
		{Id: 2, Nome: "Sapato", Preco: 19.00},
		{Id: 3, Nome: "Tenis", Preco: 30.00},
	}
	p2, _ := json.Marshal(p1)
	fmt.Println(string(p2))


	// Json para struct
	var p3 []Product

	if err := json.Unmarshal(p2, &p3); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Voltando para strcut o json tamanho da lista de produtos: ", len(p3))
	for _, p := range p3 {
		fmt.Printf("\nID: %d\tNome: %s\t%.2f", p.Id, p.Nome, p.Preco)
	}
	fmt.Println()

	byt := []byte(`{"Nome":"Hospagem", "preco": 12.99}`)
	var dat = make(map[string]interface{})
	if err := json.Unmarshal(byt, &dat); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dat)	
 
}