package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:kraken1288@localhost/project_colaboration_apps?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// insertResult, err := db.Exec(`INSERT INTO "test"(name) VALUES ('LOLOLOLOLOL')`)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(insertResult)

	rows, err := db.Query(`SELECT * FROM "test"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var names []string
	var dates []time.Time
	if rows.Next() {
		var name string
		var date time.Time
		err := rows.Scan(&name, &date)
		if err != nil {
			panic(err)
		}
		names = append(names, name)
		dates = append(dates, date)
	}
	fmt.Println(names, dates)
}
