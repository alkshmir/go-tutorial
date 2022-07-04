package main

import "fmt"

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// 関数を引数にとる関数
func CallFunction(f func()) {
	f()
}

func main() {
	f := ReturnFunc()
	f()
	CallFunction(f)
	CallFunction(func() {
		fmt.Println("mumei kansuu")
	})
}
