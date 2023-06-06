package main

import (
	"database/sql"
	"fmt"
	"go-db-sqlite/internal/domain"
	"go-db-sqlite/internal/repository/sqlite3"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	DB = sqlite3.DBConnect()
}

func main() {
	repository := sqlite3.NewRepository(DB)
	// u := domain.User{UserName: "test 2", Email: "test2@gmail.com"}
	// err := repository.InsertData(u)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// d, e := repository.SelectData()
	// if e != nil {
	// 	fmt.Println(e.Error())
	// }
	// for _, u := range d {
	// 	fmt.Println(u)
	// }

	oldUser := domain.User{UserID: 2}
	newUser := domain.User{UserName: "juarez2", Email: "juarez132@gmail.com"}
	err := repository.UpdateData(oldUser, newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dados a ser alterado")
	fmt.Printf("Nome: %s\tEmail: %s", newUser.UserName, newUser.Email)
}
