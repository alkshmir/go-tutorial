package main

import "fmt"

type Myint int // typeは型をつくる文

// 自作型のメソッド
func (mi Myint) Print() {
	fmt.Println(mi)
}

func main() {
	var mi Myint
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	//i := 100
	//fmt.Println(mi + i) // error 他のデータ型と計算はできない

	mi.Print()
}
