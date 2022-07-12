package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	name string
	age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", ".example.sql")

	defer Db.Close()

	cmd := `CREATE TABLE IF NOT EXISTS persons(
		name STRING,
		age INT
	)`
	_, err := Db.Exec(cmd)

	if err != nil {
		log.Fatalln(err)
	}
	/*
		cmd = "INSERT INTO persons (name, age) VALUES (?, ?)"
		_, err = Db.Exec(cmd, "tarou", 20) // sqlインジェクションを防ぐため?をエスケープ
		if err != nil {
			log.Fatalln(err)
		}

		cmd = "UPDATE persons SET age = ? WHERE name = ?"
		_, err = Db.Exec(cmd, 25, "tarou") // sqlインジェクションを防ぐため?をエスケープ
		if err != nil {
			log.Fatalln(err)
		}
	*/
	/*
		cmd = "INSERT INTO persons (name, age) VALUES (?, ?)"
		_, err = Db.Exec(cmd, "hanako", 19) // sqlインジェクションを防ぐため?をエスケープ
		if err != nil {
			log.Fatalln(err)
		}
	*/

	cmd = "Select * from persons where age = ?"
	row := Db.QueryRow(cmd, 25)
	var p Person
	err = row.Scan(&p.name, &p.age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No Rows")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.name, p.age)

	cmd = "select * from persons"
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		err = rows.Scan(&p.name, &p.age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p.name, p.age)
	}
}
