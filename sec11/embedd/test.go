package main

import "fmt"

type car struct {
	engine
	tire
}

type engine struct {
	Name  string
	power int
}

type tire struct {
	Name     string
	diameter int
}

func main() {
	engine1 := engine{"good engine", 300}
	tire1 := tire{"good tire", 10}
	car1 := car{engine1, tire1}
	fmt.Println(car1)
	//fmt.Println(car1.Name) // ambiguous selector car1.Nameとなってエラー
	fmt.Println(car1.engine.Name)
}
