package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ファイルが存在する場合は上書き
	f, _ := os.Create("foo.txt")

	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"), 6) //6文字目に書き込む
	f.Seek(0, os.SEEK_END)         //ファイル末尾に移動する
	f.WriteString("Yeah")

	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)

	n, err := f.Read(bs) // bsに書き込む
	fmt.Println(n)       // 読み込んだバイト数
	fmt.Println(string(bs))

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 10) // 10バイト目
	fmt.Println(nn)
	fmt.Println(string(bs2))

	// OpenFile
	// os.O_RDWR|os.O_CREATE 読み書き可能でファイルがなければ作成
	f, err = os.OpenFile("foo.txt", os.O_RDONLY, 0666) // READONLY, permission 0666
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs = make([]byte, 128)
	n, err = f.Read(bs)

	fmt.Println(n)
	fmt.Println(string(bs))
}
