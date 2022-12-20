package main

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"

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

type ErrorResponse struct {
	Message string `json:"message"`
}

var db *sql.DB

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem Vindo!")
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	registers, err := db.Query("SELECT id, title, author FROM books")
	if err != nil {
		log.Println("❌  getAllBooks ==>" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var books []Book = make([]Book, 0)
	for registers.Next() {
		var book Book
		err = registers.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			log.Println("❌  getAllBooks ==>" + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}
	err = registers.Close()
	if err != nil {
		log.Println("❌  getAllBooks ==>" + err.Error())
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newBook Book
	json.Unmarshal(body, &newBook)

	if len(newBook.Title) == 0 || len(newBook.Author) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{"Title and Author are required"})
		return
	}

	register, errRegister := db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", newBook.Title, newBook.Author)

	if errRegister != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	generateId, _ := register.LastInsertId()
	newBook.Id = int(generateId)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func searchBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	register := db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
	var book Book

	err = register.Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(book)
}

func deleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	register := db.QueryRow("SELECT id FROM books WHERE id = ?", id)
	var bookId int
	err = register.Scan(&bookId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, err = db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateBookById(w http.ResponseWriter, r *http.Request) {
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

	register := db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
	var book Book
	err = register.Scan(&book.Id, &book.Title, &book.Author)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var newBook Book
	errorJson := json.Unmarshal(body, &newBook)
	newBook.Id = id

	if errorJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.Exec(("UPDATE books SET title = ?, author = ? WHERE id = ?"), newBook.Title, newBook.Author, newBook.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newBook)
}

func configHandles(router *mux.Router) {
	router.HandleFunc("/", mainHandle)
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", searchBookById).Methods("GET")
	router.HandleFunc("/books/{id}", updateBookById).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBookById).Methods("DELETE")
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func configDB() {
	var err error
	var connectionStr string = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	db, err = sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal(errPing.Error())
	}

	fmt.Println("✔️  Connected to database")
}

func configServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)
	configHandles(router)

	fmt.Println("Servidor rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", router)) // DefaultServeMux
}

func main() {
	configDB()
	configServer()
}
