package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "ABC"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}
