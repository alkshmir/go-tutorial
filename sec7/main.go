package main

import "fmt"

// x int, y intの意味
func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

//返り血の変数を指定することができる
func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	/*
		これはエラー
		i4 := Div(9, 3)
		fmt.Println(i4)
	*/
	// これはOK
	i4, _ := Div(9, 3)
	fmt.Println(i4)

	// これはNG
	//a := Noreturn()

	//これはOK
	Noreturn()
}
