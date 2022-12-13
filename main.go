package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Id     int
	Title  string
	Author string
}

var Books []Book = []Book{
	{Id: 1, Title: "O Senhor dos Aneis", Author: "J.R.R. Tolkien"},
	{Id: 2, Title: "O Hobbit", Author: "J.R.R. Tolkien"},
	{Id: 3, Title: "O Silmarillion", Author: "J.R.R. Tolkien"},
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem Vindo!")
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Books)
}

func configHandles() {
	http.HandleFunc("/", mainHandle)
	http.HandleFunc("/books", getAllBooks)
}

func configServer() {
	configHandles()

	fmt.Println("Servidor rodando na porta 1337")
	http.ListenAndServe(":1337", nil) // DefaultServeMux
}

func main() {
	configServer()
}
