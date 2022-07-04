package main

import "fmt"

// init()はmain()よりも先に読まれて実行される。
func init() {
	fmt.Println("init")
}

// 複数のinit()を定義することも可能
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
