package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	body, err := io.ReadAll(r.Body)
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
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, errorBody := io.ReadAll(r.Body)
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

func configHandles(router *mux.Router) {
	router.HandleFunc("/", mainHandle)
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", searchBookById).Methods("GET")
	router.HandleFunc("/books/{id}", updateBookById).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBookById).Methods("DELETE")
}

func configServer() {
	router := mux.NewRouter().StrictSlash(true)
	configHandles(router)

	fmt.Println("Servidor rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", router)) // DefaultServeMux
}

func main() {
	configServer()
}
