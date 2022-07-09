package main

import (
	"log"
	"os"
)

func main() {
	// logの出力先を変更
	log.SetOutput(os.Stdout)

	// logフォーマット
	log.SetFlags(log.LstdFlags) // 標準ログフォーマット
	log.Print("Log\n")

	// マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("Log2")

	// 時刻とファイルの行番号
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Printf("Log%d\n", 3)

	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[LOG]")
	/*
		// exit stauts 1 でプログラム終了する
		log.Fatal("Log\n")
		log.Fatalln("Log2")
		log.Fatalf("Log%d\n", 3)
	*/

	/*
		// exit status 2 でプログラム終了する
		log.Panic("Log\n")
		log.Panicln("Log2")
		log.Panicf("Log%d\n", 3)
	*/

	// ファイルにログを吐く
	f, err := os.Create("test.log")
	if err != nil {
		return
	}

	log.SetOutput(f)
	log.Println("Log")
}
