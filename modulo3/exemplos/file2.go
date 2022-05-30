package exemplos

import (
	"encoding/json"
	"fmt"
	"log"
)



type product struct {
	Name string  `json:"name"`
	Price int  `json:"price"`
	Published bool `json:"published"`
}

func ShowJson() {

	p := []product{
		{
			Name: "MackBook Pro",
			Price: 1500,
			Published: true,
		}, 
		{
			Name: "MackBook Pro 2",
			Price: 1900,
			Published: true,
		}, 

	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	objetoJson := `[{"Name":"Notbook Dell", "Price": 5900, "Published": true}, {"Name":"MacBook Pro", "Price": 900, "Published": false}, {"Name":"Notbook Positivo", "Price": 3600, "Published": false}, {"Name":"Notbook LG", "Price": 900, "Published": true}]`

	var newProduto []product

	err = json.Unmarshal([]byte(objetoJson), &newProduto)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range newProduto {
		fmt.Printf("Name: %s\tPrice: %.2f\tPublished: %t\n", p.Name, float64(p.Price), p.Published)

	}

}