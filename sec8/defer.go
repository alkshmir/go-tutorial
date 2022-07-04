package main

import (
	"fmt"
	"os"
)

func TestDefer() {
	defer fmt.Println("END") // 関数終了時に実行する処理
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
func main() {
	TestDefer()

	RunDefer() // c b a 後から追加したものから実行される(stack)
	// main関数が終了したときの処理を無名関数で書く
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	// リソース開放処理
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello"))
}
