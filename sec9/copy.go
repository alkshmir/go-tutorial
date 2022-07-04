package main

import "fmt"

func main() {
	sl := []int{100, 200}
	sl2 := sl // これは参照

	sl2[0] = 1000   //元のslも更新される
	fmt.Println(sl) // [1000, 200]

	var i int = 10
	i2 := i //これはコピー
	i2 = 1000
	fmt.Println(i, i2)

	sl1 := []int{1, 2, 3, 4, 5}
	sl3 := make([]int, 5, 10)
	//copy(コピー先, コピー元)
	n := copy(sl3, sl1) // n: コピーに成功した要素数

	fmt.Println(n, sl3)
}
