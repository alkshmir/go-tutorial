package main

import (
	"fmt"
	"os"
)

// buildする必要がある
func main() {
	fmt.Println(os.Args[0]) // コマンド名
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	fmt.Printf("length=%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
