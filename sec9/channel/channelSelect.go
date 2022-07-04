package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "A"
	/*
		//ch1 <- 1
		v1 := <-ch1
		v2 := <-ch2
		fmt.Println(v1) //deadlock
		fmt.Println(v2)
	*/
	// select文はchannelに対する処理
	// switchとは異なりどのcaseが実行されるかはランダム(後続のcaseも均等に実行するため)
	select {
	case v1 := <-ch1: // ch1からの受信
		fmt.Println(v1 + 1000)
	case v2 := <-ch2: // ch2からの受信
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2 // ch3から受信した値をch4に2倍して渡す
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("recieved from ch5:", i3)
		default:
			if n > 100 {
				break L
			}
		}
		/*
			if n > 100 {
				break
			}
		*/
	}
}
