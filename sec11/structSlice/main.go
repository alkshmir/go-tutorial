package main

import "fmt"

type User struct {
	Name string
	Age  int
	//X, Y int // まとめて定義できる
}

// 構造体のスライス
type Users []*User // Userのポインタをスライスとして格納できるUsers型を新しく定義する

// これでもよい
/*
type Users struct {
	Users []*Users
}
*/
func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}
	users = append(users, &user1, &user2, &user3, &user4)

	fmt.Println(users) // メモリ番地が表示される

	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User, 0) // スライスを宣言
	users2 = append(users2, &user1, &user2)

	for _, u := range users2 {
		fmt.Println(*u)
	}
}
