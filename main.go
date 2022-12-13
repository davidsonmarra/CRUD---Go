package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books []Book = []Book{
	{Id: 1, Title: "O Senhor dos Aneis", Author: "J.R.R. Tolkien"},
	{Id: 2, Title: "O Hobbit", Author: "J.R.R. Tolkien"},
	{Id: 3, Title: "O Silmarillion", Author: "J.R.R. Tolkien"},
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem Vindo!")
}

func rootBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllBooks(w, r)
	case "POST":
		createBook(w, r)
	default:
		fmt.Fprint(w, "Método não suportado")
	}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
		return
	}

	var newBook Book
	json.Unmarshal(body, &newBook)
	newBook.Id = len(Books) + 1
	Books = append(Books, newBook)

	encoder := json.NewEncoder(w)
	encoder.Encode(newBook)
}

func configHandles() {
	http.HandleFunc("/", mainHandle)
	http.HandleFunc("/books", rootBooks)
}

func configServer() {
	configHandles()

	fmt.Println("Servidor rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", nil)) // DefaultServeMux
}

func main() {
	configServer()
}
