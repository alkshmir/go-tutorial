package main

import "fmt"

/*
// errorはinterface型
type error interface {
	Error() string
}
*/

//error
type MyError struct {
	Message string
	ErrCode int
}

// Goの標準エラーと同じインターフェースを持つ
// GoのエラーはインターフェースなのでErrorメソッドさえ持たせればどんな型もエラーにすることができます。
func (e *MyError) Error() string {
	return e.Message
}

// MyError型はerrorインターフェースを持っているのでerrorを返す関数の戻り値にそのまま使うことができます。
func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1}
}

func main() {
	err := RaiseError()
	fmt.Printf("%T\n", err)
	fmt.Println(err.Error())

	// そのままではアクセスできない
	// fmt.Println(err.Message)

	// エラーコードを取り出すときはキャストします。
	e, ok := err.(*MyError)
	if ok { // 型変換がうまくいったとき
		fmt.Println(e.ErrCode)
	}

}
