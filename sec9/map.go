package main

import "fmt"

//map
// python でいうdict

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)
	// 取り出し
	fmt.Println(m["A"])

	m2 := map[string]int{"A": 200, "B": 300}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	//代入
	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)

	//登録されていないkey
	fmt.Println(m4[3]) // 空文字が出るエラーにはならない

	s, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok) //二つ目の返り血に取り出しに成功したかどうかが入る

	// 値の更新
	m4[2] = "US"
	fmt.Println(m4)

	// 値の削除
	delete(m4, 2) // 返り値はないので消せたかはわからない
	fmt.Println(m4)

	// 要素数の取得
	fmt.Println(len(m4))
}
