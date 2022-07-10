package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("foo.txt")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))
}
