package main

import "fmt"

// implement generator using closure
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}
