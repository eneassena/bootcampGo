package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Field struct {
	ID int 
	Name string 
	Price float64
}

type Produto struct {
	field Field
}


func main() {

	byt, e := ioutil.ReadFile("product.json") 
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(string(byt))

	var prod Field


	if err := json.Unmarshal(byt, &prod); err != nil {
		log.Fatal(err)
	}

	fmt.Println(prod)



}