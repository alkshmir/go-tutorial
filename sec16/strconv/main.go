package main

import (
	"fmt"
	"strconv"
)

func main() {
	bt := true
	fmt.Println(strconv.FormatBool(bt))
	fmt.Printf("%T\n", strconv.FormatBool(bt))

	i := strconv.FormatInt(-100, 10) //10進数
	fmt.Printf("%v, %T\n", i, i)

	// 10進数ならItoaで簡易的に変換できる
	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2)

	// 文字列を真偽値に変換
	bt1, _ := strconv.ParseBool("true")
	fmt.Printf("%v, %T\n", bt1, bt1)
	bt2, _ := strconv.ParseBool("1")
	bt3, _ := strconv.ParseBool("t")
	bt4, _ := strconv.ParseBool("T")
	bt5, _ := strconv.ParseBool("TRUE")
	bt6, _ := strconv.ParseBool("True")
	fmt.Println(bt2, bt3, bt4, bt5, bt6)

	// 文字列を整数に変換
	i3, _ := strconv.ParseInt("12345", 10, 0) // 10進数、整数の精度(bit)(0の時はデフォルト値)
	fmt.Printf("%v, %T\n", i3, i3)

	// 10進数、デフォルト精度なら簡易的に変換できる
	i10, _ := strconv.Atoi("123")
	fmt.Println(i10)
}
