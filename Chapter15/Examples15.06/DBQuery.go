package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	var id int
	var name string

	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Retrieved data from db: %d %s\n", id, name)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	err = rows.Close()
	if err != nil {
		panic(err)
	}
	db.Close()
}
