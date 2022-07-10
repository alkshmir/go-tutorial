package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)

	// compile
	re1, _ := regexp.Compile("A") // 同じ正規表現を繰り返し使う場合
	match = re1.MatchString("ABC")
	fmt.Println(match)

	// MustCompile
	re2 := regexp.MustCompile("A") // エラーが発生したときに直接runtime errorを吐く
	match = re2.MatchString("ABC")
	fmt.Println(match)

	regexp.MustCompile("\\d")
	regexp.MustCompile(`\d`) // ``で囲むとエスケープを使わなくてよい

	// i 大文字小文字を区別しない
	// m マルチラインモード
	// s .が\nにマッチ
	re3 := regexp.MustCompile(`(?i-ms)abc`)
	match = re3.MatchString("ABC")
	fmt.Println(match)

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)

	match = re4.MatchString("  ABC   ")
	fmt.Println(match)

	//繰り返し
	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaaabbbbbb"))
	fmt.Println(re6.MatchString("b"))

	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`) //英数字と_を含めた3文字
	fmt.Println(re9.MatchString("abc"))
	fmt.Println(re9.MatchString("abcdefg"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`) //英数字以外
	fmt.Println(re10.MatchString("ABC"))
	fmt.Println(re10.MatchString("あ"))

	// 正規表現のグループ
	/*
		(正規表現)　グループ（順序によるキャプチャ）
		(?:正規表現)　（キャプチャされない）
		(?:P<name>正規表現)　名前付きグループ
	*/
	re11 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re11.MatchString("abcxyz"))
	fmt.Println(re11.MatchString("ABCXYZ"))
	fmt.Println(re11.MatchString("ABCxyz"))
	fmt.Println(re11.MatchString("ABCabc"))

	fs1 := re11.FindString("AAAAABCXYZZZ")
	fmt.Println(fs1)
	fs2 := re11.FindAllString("ABCXYZABCxyz", -1)
	fmt.Println(fs2)

	// グループサブマッチ
	re12 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
			0123-145-4579
			111-222-333
			556-322-118
	`
	ms := re12.FindAllStringSubmatch(s, -1)
	fmt.Println(ms)

	for _, v := range ms {
		fmt.Println(v)
	}
	fmt.Println(ms[0][0]) // マッチした全体が入る
	fmt.Println(ms[0][1]) // サブマッチした文字列が入る
	fmt.Println(ms[0][2])

	// 置換
	re13 := regexp.MustCompile(`\s+`)
	fmt.Println(re13.ReplaceAllString("AAA BBB CCC", ","))

	re14 := regexp.MustCompile(`、|。`)
	fmt.Println(re14.ReplaceAllString("こんにちは。私は、Golang。", ""))

	re15 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re15.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ>", -1)) //分割

	re16 := regexp.MustCompile(`\s+`)
	fmt.Println(re16.Split("aaaa bb    ccccc", -1))

}
