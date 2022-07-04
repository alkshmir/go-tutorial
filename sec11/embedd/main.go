package main

import "fmt"

type T struct {
	//User User // フィールド名　と型が同じでもよい
	User //こう書いてもよい こう書くとUser型のフィールドに直接アクセスできる
}

type User struct {
	Name string
	Age  int
	//X, Y int // まとめて定義できる
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	fmt.Println(t.Name)
	t.SetName()
	fmt.Println(t.Name)
}
