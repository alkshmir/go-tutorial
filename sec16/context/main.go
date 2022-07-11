package main

import (
	"context"
	"fmt"
	"time"
)

// コンテキストを使ってタイムアウト処理
func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("stop")
	ch <- "実行結果"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()

	// 一秒のタイムアウトを設定
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	go longProcess(ctx, ch)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("#####ERROR#####")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}
	}

	fmt.Println("loop ended")
}
