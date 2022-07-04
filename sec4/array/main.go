package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)

	fmt.Printf("%T\n", arr1) // [3]int が型名 要素数が違うものは違う型

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])

	arr2[2] = "C"
	fmt.Println(arr2)

	//var arr5 [4]int
	//arr5 = arr1
	//fmt.Println(arr5)

	fmt.Println(len(arr1))
}
