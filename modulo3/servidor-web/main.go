package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	response := Response{Message: "Golang Servidor"}

	data, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}

// trabalhando com servidor web
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
