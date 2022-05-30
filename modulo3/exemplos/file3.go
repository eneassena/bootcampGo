package exemplos

import (
 
	"net/http"
)


func helloHandler(w http.ResponseWriter, req *http.Request) { 
	w.Write([]byte("Ola"))
}


func Web() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}