package main

import "fmt"

func main() {
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "DE")
	fmt.Println(5 - 4)
	fmt.Println(5 * 4)
	fmt.Println(60 / 3)
	fmt.Println(9 % 4)

	n := 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)

	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(4 <= 8)
	fmt.Println(1 >= 8)

	fmt.Println("ABC" == "ABC")
	var s string
	fmt.Println("" == s)
	var t string
	fmt.Println(s == t)
}
