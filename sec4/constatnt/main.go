package main

import "fmt"

// 先頭大文字にするとほかのパッケージから読みだせる
// 先頭小文字の場合はこのパッケージだけ
const Pi = 3.14
const (
	URL      = "http://example.com"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 定数は巨大でもよい
//var Big int = 7548947098247305298752483097508392470934
//var Big = 754894709824730529875248

const (
	c0 = iota // 整数の連番を返す
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	//Pi = 3
	//fmt.Println(Pi)
	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F) // 1 1 1 A A A

	//fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2)
}
