package main

import (
	"flag"
	"fmt"
)

// flag
// 引数やオプションを処理するための仕組み
// os.argsを使わなくてよい

func main() {
	// オプションの値を格納する変数
	var (
		max int
		msg string
		x   bool
	)

	// IntVar 整数オプション
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	flag.StringVar(&msg, "m", "", "処理メッセージ") //デフォルトち空文字列
	// オプションが与えられればtrue, なければfalse
	flag.BoolVar(&x, "x", false, "拡張オプション")

	// コマンドラインをパース
	flag.Parse()

	fmt.Println("処理数の最大値 = ", max)
	fmt.Println("処理メッセージ = ", msg)
	fmt.Println("拡張オプション = ", x)

	//存在しないオプションを指定するとヘルプメッセージを自動で表示
}
