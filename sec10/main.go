package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}
func main() {
	var n int = 100
	fmt.Println(n)

	Double(n) // nがコピーされる

	// アドレス
	fmt.Println(&n)

	var p *int = &n //nのアドレスへのポインタ
	fmt.Println(p)  // アドレスが表示される
	fmt.Println(*p) // ポインタが指し示す実体(100)

	*p = 300       //実体を書き換える
	fmt.Println(n) // 300

	n = 200
	fmt.Println(*p) // 200

	Doublev2(&n) // コピーしないで渡す
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(n)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}

	Doublev3(sl)    // スライスは参照渡し
	fmt.Println(sl) // 2, 4, 6
}
