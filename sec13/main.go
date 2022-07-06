package main

import (
	"fmt"
	. "fmt" // パッケージ名の指定が不要となる(非推奨)
	f "fmt" // fmtパッケージにfという名前をつける
	"udemy/sec13/foo"
)

//scope

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	b = s
	{
		b := "BBBB" // {}内は異なる変数スコープなので再定義してよい
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	f.Println(foo.Max)
	// 参照不可
	//fmt.Println(foo.min)
	f.Println(foo.ReturnMin())

	Println(foo.Max)

	fmt.Println(appName())
	// fmt.Println(AppName) //エラー

	fmt.Println(Do("AAA"))
}
