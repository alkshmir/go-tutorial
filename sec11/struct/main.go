package main

import "fmt"

// 構造体：クラスっぽいやつ

type User struct {
	Name string
	Age  int
	//X, Y int // まとめて定義できる
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	fmt.Println(user1.Age)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{} //暗黙的宣言
	fmt.Println(user2)
	user2.Name = "user2"
	user2.Age = 20

	user3 := User{Name: "user3", Age: 30} // 初期値
	fmt.Println(user3)

	user4 := User{"user4", 40} // フィールドを指定しない(定義した順に入る)
	fmt.Println(user4)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User) // 構造体に対するポインタを返す
	fmt.Println(user7)

	user8 := &User{} // user7とおなじ
	fmt.Println(user8)
	//関数の引数として構造体を渡すときに使いがち

	UpdateUser(user1)  // 更新されない
	UpdateUser2(user8) // 更新される

	fmt.Println(user1)
	fmt.Println(user8)

	UpdateUser2(&user1) //更新される
	fmt.Println(user1)
}
