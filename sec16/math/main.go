package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(15))

	// 絶対値
	fmt.Println(math.Abs(-100))

	//累乗
	fmt.Println(math.Pow(0, 2))

	// ３乗根
	fmt.Println(math.Cbrt(8))

	// 最大値、最小値
	fmt.Println(math.Max(1, 2)) //2つだけ
	fmt.Println(math.Min(1, 2))

	// 小数点以下を切り捨てる
	fmt.Println(math.Trunc(3.14))

	// math.Floor
	// math.Ceil

}
