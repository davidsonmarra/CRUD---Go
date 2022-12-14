package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) == 2 || len(parts) == 3 && parts[2] == "" {
		switch r.Method {
		case "GET":
			getAllBooks(w, r)
		case "POST":
			createBook(w, r)
		default:
			fmt.Fprint(w, "Método não suportado")
		}
	} else if len(parts) == 3 || len(parts) == 4 && parts[3] == "" {
		switch r.Method {
		case "GET":
			searchBookById(w, r)
		case "DELETE":
			deleteBookById(w, r)
		case "PUT":
			updateBookById(w, r)
		default:
			fmt.Fprint(w, "Método não suportado")
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
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

func searchBookById(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	var id, err = strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, book := range Books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteBookById(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[2])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, book := range Books {
		if book.Id == id {
			left := Books[:i]
			right := Books[i+1:]
			Books = append(left, right...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateBookById(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[2])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, errorBody := ioutil.ReadAll(r.Body)
	if errorBody != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newBook Book
	errorJson := json.Unmarshal(body, &newBook)
	newBook.Id = id

	if errorJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, book := range Books {
		if book.Id == id {
			Books[i] = newBook
			json.NewEncoder(w).Encode(newBook)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func configHandles() {
	http.HandleFunc("/", mainHandle)
	http.HandleFunc("/books", rootBooks)

	// e.g. GET /books/{id}
	http.HandleFunc("/books/", rootBooks)
}

func configServer() {
	configHandles()

	fmt.Println("Servidor rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", nil)) // DefaultServeMux
}

func main() {
	configServer()
}
