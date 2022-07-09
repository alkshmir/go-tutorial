package main

import (
	"fmt"
	"strings"
)

func main() {
	// セパレータを指定して文字列を結合する
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	//　部分文字列を検索する
	i1 := strings.Index("ABCDEABCDE", "A") //最初の番号
	fmt.Println(i1)

	i2 := strings.LastIndex("ABCDABCD", "BC") //最後の番号
	fmt.Println(i2)

	i3 := strings.IndexAny("ABC", "ABC") // 第二引数の文字のうちどれかが含まれるインデックス位置
	i4 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i3, i4)

	b1 := strings.HasPrefix("ABC", "A") // pythonのstartwithみたいなやつ
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2)

	b3 := strings.Contains("ABC", "B")
	b4 := strings.ContainsAny("ABC", "BD")
	fmt.Println(b3, b4)

	i6 := strings.Count("ABCABC", "B")
	i7 := strings.Count("ABCABC", "") // 空文字の場合は文字列の長さ+1
	fmt.Println(i6, i7)

	// 文字列を繰り返して結合する
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4)

	// 文字列の置換
	s5 := strings.Replace("AAABBBCCC", "A", "B", 1) // 最後は置換する回数の最大数 -1 は無限
	s6 := strings.Replace("AAAAAAA", "A", "B", -1)
	fmt.Println(s5, s6)

	// 文字列の分割
	s7 := strings.Split("A.B.C.D.E", ".")
	fmt.Println(s7)
	s8 := strings.SplitAfter("A.B.C.D.E", ".")
	fmt.Println(s8)
	s9 := strings.SplitAfterN("A.B.C.D.E", ".", 2)
	fmt.Println(s9)
	s10 := strings.SplitN("A.B.C.D.E", ".", 2)
	fmt.Println(s10)

	// 大文字小文字の変換
	s11 := strings.ToLower("ABC")
	s13 := strings.ToUpper("abc")
	fmt.Println(s11, s13)

	// 文字列から空白を取り除く
	s15 := strings.TrimSpace("    -    ABC -    - ")
	// 全角スペースも取り除く
	s16 := strings.TrimSpace("　　　ABC　　　")
	fmt.Println(s15, s16)

	// スペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[1])
}
