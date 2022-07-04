package main

import "fmt"

// 「stringを返す関数」を返す関数
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("Name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))
	// クロージャーを使うとstore変数を内部状態として引き継いでいる
	// 普通に関数定義するとstore変数は関数がreturnしたときに破棄される。
}
