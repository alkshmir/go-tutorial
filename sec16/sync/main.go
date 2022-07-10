package main

import (
	"fmt"
	"sync"
	"time"
)

//ミューテックスによる同期処理

// レースコンディション：並列動作する複数のスレッドが同一リソースに同時にアクセスしたときに起こる問題のこと

//ミューテックスでロックすることで同時に1つのgo routineのみがリソースにアクセスできるようにする
var mutex *sync.Mutex

var st struct{ A, B, C int }

/* 並行処理すると一貫性が崩れる
func UpdateAndPrint(n int) {
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)
}
*/
func UpdateAndPrint(n int) {
	//ロック
	mutex.Lock()

	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)

	//アンロック
	mutex.Unlock()
}
func main() {
	//初期化
	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	for {

	}
}
