package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Configura o roteamento do servidor
	http.HandleFunc("/hello", helloHandler)

	// Inicia o servidor na porta 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
