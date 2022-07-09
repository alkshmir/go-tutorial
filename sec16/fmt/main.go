package main

import (
	"fmt"
	"os"
)

func main() {
	s := fmt.Sprint("Hello")
	s1 := fmt.Sprintf("%v\n", "Hello")
	s2 := fmt.Sprintln("Hello")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)

	// Fprint 書き込み先指定
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")
}
