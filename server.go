package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/configmap", configMap)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}

func configMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "Hello, %s.", string(data))
}
