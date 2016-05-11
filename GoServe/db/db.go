package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"go:go query@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT * FROM Players")
	defer rows.Close()
	if err != nil {
		fmt.Println(err, "ruh roh")
	}
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			fmt.Print(err)
			break
		}
		for _, c := range columns {
			fmt.Println(rows.Scan(&c))
		}
	}
	defer db.Close()
}
