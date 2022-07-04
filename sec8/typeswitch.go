package main

import "fmt"

func anything(a interface{}) {
	//fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}
func main() {
	// 型assertion 動的に型チェック
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i := x.(int)       // interface型変数.(変換したい型)
	fmt.Println(i + 2) // 5
	fmt.Println(x)     // 更新されない
	// これはエラー
	//fmt.Println(x + 2)

	//これはruntime error
	//f := x.(float64)

	// これはOK
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64) // 変換が失敗するとfalse

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}

	// 型スイッチ
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	// 復元した値も使う場合
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}
