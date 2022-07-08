package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(256)
	fmt.Println(rand.Float64()) // 0.0 ~ 1.0
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	// 現在時刻をSeedにする
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100)) // 0 ~ 99
	fmt.Println(rand.Int())
	fmt.Println(rand.Int31()) // Int32型
	fmt.Println(rand.Int63())
	fmt.Println(rand.Uint32()) // Uintは32なので注意
}
