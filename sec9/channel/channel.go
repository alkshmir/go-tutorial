package main

import "fmt"

func main() {
	// goroutine間でデータの送受信を行うデータ構造(queue)
	var ch1 chan int

	// 受信専用チャネル
	//var ch2 <-chan int

	// 送信専用チャネル
	//var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	//チャネルの容量(buffer size)
	fmt.Println(cap(ch1)) // 0
	fmt.Println(cap(ch2)) // 0

	// buffer sizeを指定して生成
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3)) // 5

	fmt.Println(len(ch3))
	// データをチャネルに送信
	ch3 <- 1
	// チャネル内の要素数
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	//チャネルからデータを受信する
	i := <-ch3
	fmt.Println(i) // FIFO
	i2 := <-ch3
	fmt.Println(i2) // FIFO
	fmt.Println("len", len(ch3))

	//変数に入れなくてもよい
	fmt.Println(<-ch3)

	// バッファサイズを超えると
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6 //deadlock
}
