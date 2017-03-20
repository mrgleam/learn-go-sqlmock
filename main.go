package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type user struct {
	name string
}

func getNameByAge(age int, db *sql.DB) (u []*user, err error) {
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	defer rows.Close()

	var users []*user

	for rows.Next() {
		p := &user{}
		if err := rows.Scan(&p.name); err != nil {
			return users, err
		}
		users = append(users, p)
	}
	return users, err
}

func main() {
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	getNameByAge(21, db)
}
