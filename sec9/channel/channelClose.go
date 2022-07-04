package main

import (
	"fmt"
	"time"
)

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break // 受信しなくなったら終了する
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	/*
		ch1 <- 1
		close(ch1)
		// closeされたchannelに対してデータを送信するとruntime error
		//ch1 <- 1
		// 受信はできる
		//fmt.Println(<-ch1)

		i, ok := <-ch1
		fmt.Println(i, ok) // ok はchannelのopen状態を表す。closedかつチャネルのバッファが空の時false

		i2, ok := <-ch1
		fmt.Println(i2, ok) // closedかつチャネルのバッファが空の時false
	*/

	go reciever("1.goroutine", ch1)
	go reciever("2.goroutine", ch1)
	go reciever("3.goroutine", ch1)

	i := 0
	for i < 10 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(1 * time.Second) // goroutineの完了を待つ(本当はsyncパッケージなどを使う)
}
