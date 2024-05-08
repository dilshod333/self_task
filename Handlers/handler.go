package Handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// create book
type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "Dilshod@2005"
	port     = 5432
	dbname   = "dilshod"
)

func Connection() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreatBook(w http.ResponseWriter, r *http.Request) {
	var book = &Book{
		Title:  "Can't hurt m",
		Author: "David Goggins",
	}
	db := Connection()
	_, err := db.Exec("insert into use(title, author) values($1, $2)", book.Title, book.Author)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Successfully created....200")

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	db := Connection()

	rows, err := db.Query("select * from use")

	if err != nil {
		log.Fatal(err)
	}
	var book []Book
	for rows.Next() {

		var bookList Book
		if err := rows.Scan(&bookList.Id, &bookList.Title, &bookList.Author); err != nil {
			log.Fatal(err)
		}
		book = append(book, bookList)

	}

	w.Header().Set("content-type", "application/json") // 
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	book.Id = 1
	book.Title = "Dosonbek"
	book.Author = "AuthorDostonbek"

	db := Connection()
	defer db.Close()

	_, err := db.Exec("UPDATE use SET title=$1, author=$2 WHERE id=$3", book.Title, book.Author, book.Id)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		log.Println("Failed to update book:", err)
		return
	}

	fmt.Fprintf(w, "Book updated successfully")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	book.Id = 3

	db := Connection()

	_, err := db.Exec("delete from use where id=$1", book.Id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "It is deleted brother look.... /get route")
}
