package main

import "fmt"

func main() {
	var sl []int //整数型スライス
	//var arr [3]int //整数型配列；要素数が必要
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5) //要素数5のスライスを作成
	fmt.Println(sl4)

	// 値の更新
	sl2[0] = 1000
	fmt.Println(sl2)

	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])
	// これはエラー
	//fmt.Println(sl5[-1])
	// 最後の要素はこうする
	fmt.Println(sl5[len(sl5)-1])
	fmt.Println(sl5[1 : len(sl5)-1])
}
