package main

import "fmt"

type User struct {
	Name string
	Age  int
	//X, Y int // まとめて定義できる
}

// レシーバー
// クラスで言うメソッドっぽいもの
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

// レシーバー側で参照を受けるか値を受けるかが変わる
// 基本的にポインタ型にすべき
func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("A") // 更新されない
	user1.SayName()

	user1.SetName2("B") // 呼び出し側はポインタ型で定義しなくてよい
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.SetName("C") // 更新されない
	user2.SayName()
	user2.SetName2("C") // 更新される
	user2.SayName()
}
