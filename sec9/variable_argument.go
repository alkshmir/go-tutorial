package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	// sがスライスとなって渡される
	for _, v := range s {
		n += v
	}
	return n
}
func main() {
	fmt.Println(Sum(1, 2, 3))

	fmt.Println(Sum()) // 0

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
}
