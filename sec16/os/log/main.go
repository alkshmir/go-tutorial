package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err) // exit status 1
	}
}
