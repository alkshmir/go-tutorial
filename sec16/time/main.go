package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // 現在時刻
	fmt.Println(t)

	t2 := time.Date(2022, 7, 6, 10, 10, 10, 0, time.Local)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay()) //元日からの経過日
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())

	//time.Duration型
	fmt.Println(time.Hour)

	// time.Duration型を文字列から生成する
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	t3 := t.Add(2*time.Minute + 15*time.Second)
	//現在時刻の2分30秒後
	fmt.Println(t3)

	//差分
	d2 := t.Sub(t2) // t - t2
	fmt.Println(d2)

	//時刻を比較する
	fmt.Println(t.Before(t2))
	fmt.Println(t.After(t2))

	// sleep
	time.Sleep(5 * time.Second)
}
