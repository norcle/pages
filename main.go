package main

import (
	conn "highload/connection"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn.ConnectDb()
}
