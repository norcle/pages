package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var DB *sql.DB

func main() {
	ConnectDb()

	r := mux.NewRouter()
	r.HandleFunc("/{slug}", pageHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
	defer DB.Close()
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["slug"])
	body := findPage(vars["slug"])
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}

type Page struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Body string `json:"body"`
}

func findPage(slug string) string {
	var page Page
	var err error
	err = DB.QueryRow("SELECT slug, body FROM pages WHERE slug = ?", slug).Scan(&page.Slug, &page.Body)
	if err != nil {
		log.Println(err.Error()) // proper error handling instead of panic in your app
	}
	log.Println(page.Body)
	return page.Body
}

func ConnectDb() {
	fmt.Println("Go MySQL Tutorial")
	var err error
	DB, err = sql.Open("mysql", "main:qwerty@tcp(127.0.0.1:3306)/highload")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected")
	//migration()
}

func migration() {
	createPageQuery, err := DB.Query("CREATE TABLE IF NOT EXISTS pages (id  PRIMARY KEY, slug varchar(255), body text)")
	if err != nil {
		log.Println(err.Error())
	}
	defer createPageQuery.Close()
}
