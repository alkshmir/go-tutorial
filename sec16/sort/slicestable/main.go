package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

func main() {
	el := []Entry{
		{"A", 20},
		{"F", 50},
		{"i", 30},
		{"b", 10},
		{"y", 30},
		{"c", 30},
	}

	fmt.Println(el)

	// 名前を昇順にソート
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	// 名前順にソートした結果は無視されてValue順になる
	fmt.Println(el)

	// 安定ソート
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	//
	fmt.Println(el)
}
