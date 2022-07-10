package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//構造体からJSONへの変換

type A struct{}

type User struct {
	ID      int       `json:"id"`    // jsonに変換した際の表現を変更できる文字列に変えたい場合はjson:"id,string"
	Name    string    `json:"name"`  // 表示したくない場合は"-"
	Email   string    `json:"email"` // 空（型の初期値）の場合表示したくないとき"omitempty"
	Created time.Time `json:"created"`
	A       A         `json:"A,omitempty"`
}

// Marshalのカスタム
// MarshalJSONというメソッドを定義するとjson.Marshal()が呼ばれたときに実行される
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

//Unmarshalのカスタム
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Created = time.Now()

	// Marshal JSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs) //[]uint8

	u2 := new(User)

	// Unmarshal Jsonを構造体に変換
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)
}
