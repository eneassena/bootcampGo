package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)



type Alunos struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,uppercase"`
	Idade string `json:"idade" validate:"required,number"`
 } 
 
// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() { 
	a := Alunos{Name: "Carlos",Email: "carlos@hotmail.com", Idade: "14"}	
	validate = validator.New()
	err := validate.Struct(a)
	if err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			fmt.Println(err)
		}
		fmt.Println(err)
	}
}