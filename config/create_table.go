package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("-- CREATE TABLE --")

	// open connection mysql
	fmt.Println("Connecting to database...")
	var connectionStr string = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", connectionStr)

	if err != nil {
		log.Fatal(err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal(err.Error())
	}

	// create table
	fmt.Println("Creating table...")
	_, errCreate := db.Exec(
		"CREATE TABLE books (" +
			"id INT NOT NULL AUTO_INCREMENT," +
			"title VARCHAR(100) NOT NULL," +
			"author VARCHAR(50) NOT NULL," +
			"PRIMARY KEY (id)" +
			")",
	)

	if errCreate != nil {
		log.Fatal(errCreate.Error())
	}

	// create registers
	fmt.Println("Creating registers...")

	_, errInset := db.Exec(
		"INSERT INTO books (title, author) VALUES " +
			"('O Senhor dos Aneis', 'J.R.R. Tolkien')," +
			"('O Silmarillion', 'J.R.R. Tolkien')," +
			"('O Hobbit', 'J.R.R. Tolkien')")

	if errInset != nil {
		log.Fatal(errInset.Error())
	}

	fmt.Println("END")
}
