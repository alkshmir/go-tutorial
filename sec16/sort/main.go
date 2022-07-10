package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 7, 8, 1, 9, 4, 10, 5}
	s := []string{"A", "x", "a"}

	fmt.Println(i, s)
	sort.Ints(i)
	sort.Strings(s)
	fmt.Println(i, s)
}
