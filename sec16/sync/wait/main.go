package main

import (
	"fmt"
	"sync"
)

// sync.WaitGroup

func main() {
	// WaitGroupを生成
	wg := new(sync.WaitGroup)
	//待ち受けるGoroutineの数
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() // 完了を通知する
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done() // 完了を通知する
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done() // 完了を通知する

	}()
	wg.Wait() // wg.Addで指定した回数wg.Doneが実行されるまで待つ

}
