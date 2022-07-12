package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	name string
	age  int
}

func main() {
	Db, err := sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable port=5434")

	if err != nil {
		log.Fatalln(err)
	}
	defer Db.Close()
	cmd := `CREATE TABLE IF NOT EXISTS persons(
		name varchar(255),
		age integer
	)`
	_, err = Db.Exec(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	// C
	/*
		cmd = "INSERT INTO persons (name, age) VALUES ($1, $2)" // postgresは?でなく$
		_, err = Db.Exec(cmd, "Nancy", 20)                      // sqlインジェクションを防ぐため$をエスケープ
		if err != nil {
			log.Fatalln(err)
		}
	*/
	// R
	cmd = "Select * from persons where age = $1"
	row := Db.QueryRow(cmd, 20)
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

	// U
	/*
		cmd = "UPDATE persons SET age = $1 WHERE name = $2"
		_, err = Db.Exec(cmd, 25, "Nancy") // sqlインジェクションを防ぐため$をエスケープ
		if err != nil {
			log.Fatalln(err)
		}
	*/
	// D
	cmd = "DELETE from persons where name = $1"
	_, err = Db.Exec(cmd, "Nancy") // sqlインジェクションを防ぐため$をエスケープ
	if err != nil {
		log.Fatalln(err)
	}
}
