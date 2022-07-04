package main

import "fmt"

func main() {
	// panic recover
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Start")
}
