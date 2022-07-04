package main

import "fmt"

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	// 条件式の前に;区切りで簡易文を付けられる
	// ifブロックの外側からはアクセスできない
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}
	//これはエラー
	//fmt.Println(b)

	x := 0
	if x := 2; true {
		fmt.Println(x) //2
	}
	fmt.Println(x) // 0
}
