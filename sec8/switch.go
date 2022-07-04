package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 1, 2: //1か2
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	// nはswitch 文内のみで有効
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	// 論理式を評価出来る
	n = 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	//これはエラー
	//case n > 3 && n < 6:
	//	fmt.Println("3 < n < 6")
	default:
		fmt.Println("I don't know")
	}
}
