package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	// MD5
	h := md5.New()
	io.WriteString(h, "ABCDE") // hにABCDEを書き込む

	fmt.Println(h.Sum(nil)) //h.Sum(nil)でハッシュ値のバイト配列を返す

	fmt.Printf("%x\n", h.Sum(nil)) //%xで16進数文字列に変換

	// sha-256
	s256 := sha256.New()
	io.WriteString(s256, "ABCDE")
	fmt.Printf("%x\n", s256.Sum(nil))

	// sha-512
	s512 := sha512.New()
	io.WriteString(s512, "ABCED")
	fmt.Printf("%x\n", s512.Sum(nil))
}
