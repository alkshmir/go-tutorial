package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	// goroutine
	// 次の処理に入れない
	//sub()

	// goを入れると並行処理可能
	go sub()
	go sub()
	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
