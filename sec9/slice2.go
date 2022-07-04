package main

import "fmt"

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	// スライスに要素を追加する
	sl = append(sl, 300)
	fmt.Println(sl)

	// 一度に複数の値を追加できる
	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	//make
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	// len
	fmt.Println(len(sl2))

	// capacity
	// 内部的に持っている配列の長さ
	fmt.Println(cap(sl2))

	// capacityを指定してmake
	sl3 := make([]int, 5, 10)
	// len
	fmt.Println(len(sl3))

	// capacity
	fmt.Println(cap(sl3)) // 10

	// capacityを超えてしまうので、新たに基底配列をアロケートして古い基底配列からコピーする
	// 新しいcapacityは現在の倍になるように設定される
	sl3 = append(sl3, 1, 2, 3, 5, 6, 7)
	// len
	fmt.Println(len(sl3)) // 11

	// capacity
	// 前のcapacityの倍になるように設定される
	fmt.Println(cap(sl3)) // 20

}
