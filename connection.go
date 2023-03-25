package connection

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectDb() {
	fmt.Println("Go MySQL Tutorial")
	var err error
	DB, err = sql.Open("mysql", "main:qwerty@tcp(127.0.0.1:3306)/highload")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected")
	migration()

	defer DB.Close()
}

func migration() {
	createPageQuery, err := DB.Query("CREATE TABLE IF NOT EXISTS pages (slug varchar(255) index, body text)")
	if err != nil {
		panic(err.Error())
	}
	defer createPageQuery.Close()
}
