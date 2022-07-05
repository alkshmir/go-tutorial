package main

import "fmt"

//fmtパッケージのstringer型はインターフェース
//文字列変換したのときの表現方法を定義する
/*
type Stringer interface {
	String() string
}
*/

type Point struct {
	A int
	B string
}

// String()メソッドを定義すればfmt.Printlnしたときの表現方法を定義できる
func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)

}
