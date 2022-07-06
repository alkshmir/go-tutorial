package main

import (
	"fmt"
	"os"
)

// os

func main() {
	//os.Exit(1) // 異常終了

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0) // defer文は実行されない
}
