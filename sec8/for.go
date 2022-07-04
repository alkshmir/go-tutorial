package main

import "fmt"

func main() {
	// 無限ループ
	//for {
	//	fmt.Println("Loop")
	//}

	// while文はなくfor文のみ
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	// 普通のfor文
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 配列に対するfor文
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

	}

	// indexを取得
	for i, v := range arr {
		fmt.Println(i, v)
	}
	// indexのみを取り出すとき
	for i := range arr {
		fmt.Println(i)
	}

	//slice
	sl := []string{"Python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	//map
	// pythonのdictみたいな感じ
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// keyだけ取得したい場合
	for k := range m {
		fmt.Println(k)
	}
}
